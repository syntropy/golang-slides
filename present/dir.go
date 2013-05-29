// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"

	"github.com/syntropy/golang-slides/pkg/present"
)

func init() {
	http.HandleFunc("/", dirHandler)
}

// extensions maps the presentable file extensions to the name of the
// template to be executed.
var extensions = map[string]RenderFunc{
	".ioslide": renderIOSlide,
	".goslide": renderGoSlide,
}

// dirHandler serves a directory listing for the requested path, rooted at basePath.
func dirHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		http.Error(w, "not found", 404)
		return
	}
	const base = "."
	name := filepath.Join(base, r.URL.Path)
	if isDoc(name) {
		err := extensions[filepath.Ext(name)](w, basePath, name)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
		}
		return
	}
	if isDir, err := dirList(w, name); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	} else if isDir {
		return
	}
	http.FileServer(http.Dir(base)).ServeHTTP(w, r)
}

type RenderFunc func(w io.Writer, base, docFile string) error

func isDoc(path string) bool {
	_, ok := extensions[filepath.Ext(path)]
	return ok
}

type IOSlides struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Date     string `json:"date"`
	Hashtag  string `json:"hashtag"`

	Author   string      `json:"author"`
	Work     string      `json:"work"`
	Email    string      `json:"email"`
	Twitter  string      `json:"twitter"`
	GPlus    string      `json:"gplus"`
	Homepage string      `json:"homepage"`
	HTML     interface{} `json:"-"`
}

func renderIOSlide(w io.Writer, base, docFile string) error {
	f, err := os.Open(docFile)
	if err != nil {
		return err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	var p IOSlides
	err = dec.Decode(&p)
	if err != nil {
		return err
	}
	rest, err := ioutil.ReadAll(io.MultiReader(dec.Buffered(), f))
	if err != nil {
		return err
	}
	p.HTML = template.HTML(string(rest))

	tmplPath := filepath.Join(base, "templates/ioslide/slide.tmpl")
	tmpl := template.Must(template.ParseFiles(tmplPath))
	err = tmpl.Execute(w, p)

	if err != nil {
		return err
	}
	return nil
}

func renderGoSlide(w io.Writer, base, docFile string) error {
	f, err := os.Open(docFile)
	if err != nil {
		return err
	}
	defer f.Close()
	doc, err := present.Parse(f, docFile, 0)
	if err != nil {
		return err
	}

	// Locate the template file.
	actionTmpl := filepath.Join(base, "templates/goslide/action.tmpl")
	contentTmpl := filepath.Join(base, "templates/goslide/slides.tmpl")

	// Read and parse the input.
	tmpl := present.Template()
	if _, err := tmpl.ParseFiles(actionTmpl, contentTmpl); err != nil {
		return err
	}

	// Execute the template.
	return doc.Render(w, tmpl)
}

// dirList scans the given path and writes a directory listing to w.
// It parses the first part of each .slide file it encounters to display the
// presentation title in the listing.
// If the given path is not a directory, it returns (isDir == false, err == nil)
// and writes nothing to w.
func dirList(w io.Writer, name string) (isDir bool, err error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return false, err
	}
	if isDir = fi.IsDir(); !isDir {
		return false, nil
	}
	fis, err := f.Readdir(0)
	if err != nil {
		return false, err
	}
	d := &dirListData{Path: name}
	for _, fi := range fis {
		// skip the pkg directory
		if name == "." && fi.Name() == "pkg" {
			continue
		}
		e := dirEntry{
			Name: fi.Name(),
			Path: filepath.Join(name, fi.Name()),
		}
		if fi.IsDir() && showDir(e.Name) {
			d.Dirs = append(d.Dirs, e)
			continue
		}
		if isDoc(e.Name) {
			d.Slides = append(d.Slides, e)
		}
	}
	if d.Path == "." {
		d.Path = ""
	}
	sort.Sort(d.Dirs)
	sort.Sort(d.Slides)
	sort.Sort(d.Articles)
	sort.Sort(d.Other)
	return true, dirListTemplate.Execute(w, d)
}

// showFile returns whether the given file should be displayed in the list.
func showFile(n string) bool {
	switch filepath.Ext(n) {
	case ".pdf":
	case ".html":
	case ".go":
	default:
		return isDoc(n)
	}
	return true
}

// showDir returns whether the given directory should be displayed in the list.
func showDir(n string) bool {
	if len(n) > 0 && (n[0] == '.' || n[0] == '_') || n == "present" {
		return false
	}
	return true
}

type dirListData struct {
	Path                          string
	Dirs, Slides, Articles, Other dirEntrySlice
}

type dirEntry struct {
	Name, Path, Title string
}

type dirEntrySlice []dirEntry

func (s dirEntrySlice) Len() int           { return len(s) }
func (s dirEntrySlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s dirEntrySlice) Less(i, j int) bool { return s[i].Name < s[j].Name }

var dirListTemplate = template.Must(template.New("").Parse(dirListHTML))

const dirListHTML = `<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <title>Talks - The Go Programming Language</title>
  <link type="text/css" rel="stylesheet" href="/static/goslide/dir.css">
  <script src="/static/goslide/dir.js"></script>
</head>
<body>


<div id="page">

  <h1>Talks</h1>

  {{with .Path}}<h2>{{.}}</h2>{{end}}

  {{with .Articles}}
  <h4>Articles:</h4>
  <dl>
  {{range .}}
  <dd><a href="/{{.Path}}">{{.Name}}</a>: {{.Title}}</dd>
  {{end}}
  </dl>
  {{end}}

  {{with .Slides}}
  <h4>Slide decks:</h4>
  <dl>
  {{range .}}
  <dd><a href="/{{.Path}}">{{.Name}}</a>: {{.Title}}</dd>
  {{end}}
  </dl>
  {{end}}

</div>

</body>
</html>`
