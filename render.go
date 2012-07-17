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

	slides := []Slide{}

	if content, err := ioutil.ReadFile(flag.Arg(0)); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	} else {
		for _, name := range strings.Split(string(content), "\n") {
			if name == "" {
				continue
			}
			if slide, err := ParseSlide(*slidedir, name); err != nil {
				fmt.Fprintf(os.Stderr, "Error while parsing slide %s: %v\n", name, err)
				os.Exit(1)
			} else {
				slides = append(slides, slide...)
			}
		}
	}

	if err := RenderSlides(out, slides, *template); err != nil {
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
					switch {
						case fields[0] == "Title":
							slide.Title = strings.TrimSpace(fields[1])
						case fields[0] == "Class":
							slide.Class = strings.TrimSpace(fields[1])
						case fields[0] == "H1":
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

type TemplateArgs struct {
	Slides []Slide
}

func RenderSlides(w io.WriteCloser, slides []Slide, template_file string) error {
	tmpl, err := template.ParseFiles(template_file)
	if err == nil {
		tmpl.Execute(w, TemplateArgs{Slides: slides})
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
