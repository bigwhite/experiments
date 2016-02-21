package main

import (
	"log"
	"os"
	"text/template"
)

var items = []string{"one", "two", "three"}

func tmplbefore15() {
	var t = template.Must(template.New("tmpl").Parse(`
	<ul>
	{{range . }}
	    <li>{{.}}</li>
	{{end }}
	</ul>
	`))

	err := t.Execute(os.Stdout, items)
	if err != nil {
		log.Println("executing template:", err)
	}
}

func tmplaftergo16() {
	var t = template.Must(template.New("tmpl").Parse(`
	<ul>
	{{range . -}}
	    <li>{{.}}</li>
	{{end -}}
	</ul>
	`))

	err := t.Execute(os.Stdout, items)
	if err != nil {
		log.Println("executing template:", err)
	}
}

func main() {
	tmplbefore15()
	tmplaftergo16()
}
