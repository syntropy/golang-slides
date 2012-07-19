package main

import (
	"github.com/knieriem/markdown"
	"text/template"
	"flag"
	"os"
	"fmt"
	"io"
	"io/ioutil"
	"bufio"
	"strings"
	"bytes"
)

type Presentation struct {
	Slides    []Slide
	Title     string
	Subtitle  string
	Presenter string
}

type Slide struct {
	Id      string
	H1      string
	Title   string
	Class   string
	Content string
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [-o <outfile>] [-t <template>] [-S <slidedir>] <infile>\n", os.Args[0])
	os.Exit(1)
}


func main() {
	slidedir := flag.String("S", "slides", "slide directory")
	template := flag.String("t", "template.html", "HTML template")
	outfile := flag.String("o", "", "output file")

	flag.Parse()

	if len(flag.Args()) < 1 {
		usage()
	}

	out := os.Stdout

	if *outfile != "" {
		if f, err := os.OpenFile(*outfile, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "opening output failed: %v\n", err)
			os.Exit(1)
		} else {
			out = f
		}
	}

	pres := &Presentation{Slides: []Slide{}}

	if content, err := ioutil.ReadFile(flag.Arg(0)); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	} else {
		parts := strings.SplitN(string(content), "\n\n", 2)
		body := ""
		if len(parts) > 1 {
			for _, header := range strings.Split(parts[0], "\n") {
				fields := strings.SplitN(header, ":", 2)
				if len(fields)==2 {
					switch strings.ToLower(fields[0]) {
					case "title":
						pres.Title = strings.TrimSpace(fields[1])
					case "subtitle":
						pres.Subtitle = strings.TrimSpace(fields[1])
					case "presenter":
						pres.Presenter = strings.TrimSpace(fields[1])
					}
				}
			}
			body = parts[1]
		} else {
			body = parts[0]
		}
		for _, name := range strings.Split(body, "\n") {
			if name == "" {
				continue
			}
			if slide, err := ParseSlide(*slidedir, name); err != nil {
				fmt.Fprintf(os.Stderr, "Error while parsing slide %s: %v\n", name, err)
				os.Exit(1)
			} else {
				pres.Slides = append(pres.Slides, slide...)
			}
		}
	}

	if err := RenderSlides(out, pres, *template); err != nil {
		fmt.Fprintf(os.Stderr, "Error while rendering slides: %v\n", err)
		os.Exit(1)
	}
}

func ParseSlide(slidedir string, name string) ([]Slide, error) {
	content, err := ioutil.ReadFile(slidedir + "/" + name + ".md")
	if err != nil {
		return nil, err
	}

	slides := []Slide{}

	for _, slide_content := range strings.Split(string(content), "\n---\n") {
		slide := Slide{Id: name}

		sections := strings.SplitN(slide_content, "\n\n", 2)

		if len(sections) > 1 {
			for _, header := range strings.Split(sections[0], "\n") {
				fields := strings.SplitN(header, ":", 2)
				if len(fields) == 2 {
					switch strings.ToLower(fields[0]) {
						case "title":
							slide.Title = strings.TrimSpace(fields[1])
						case "class":
							slide.Class = strings.TrimSpace(fields[1])
						case "h1":
							slide.H1 = strings.TrimSpace(fields[1])
					}
				}
			}
			slide.Content = RenderMarkdown(sections[1])
		} else {
			slide.Content = RenderMarkdown(sections[0])
		}

		slides = append(slides, slide)
	}

	return slides, nil

}


func RenderSlides(w io.WriteCloser, pres *Presentation, template_file string) error {
	tmpl, err := template.ParseFiles(template_file)
	if err == nil {
		tmpl.Execute(w, pres)
	}
	return err
}

func RenderMarkdown(md string) string {
	htmlbuf := bytes.NewBufferString("")
	p := markdown.NewParser(&markdown.Extensions{Smart: true, FilterHTML: false})
	w := bufio.NewWriter(htmlbuf)
	p.Markdown(bytes.NewBufferString(md), markdown.ToHTML(w))
	w.Flush()
	return htmlbuf.String()
}
