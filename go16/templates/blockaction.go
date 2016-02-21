package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var items = []string{"one", "two", "three"}

var overrideItemList = `
{{define "list" -}}
	<ul>
	{{range . -}}
		<li>{{.}}</li>
	{{end -}}
	</ul>
{{end}} 
`

var tmpl = `
    Items:
	{{block "list" . -}}
	<ul>
	{{range . }}
		<li>{{.}}</li>
	{{end }}
	</ul>
	{{end}} 
`

var t *template.Template

func init() {
	t = template.Must(template.New("tmpl").Parse(tmpl))
}

func tmplBeforeOverride() {
	err := t.Execute(os.Stdout, items)
	if err != nil {
		log.Println("executing template:", err)
	}
}

func tmplafterOverride() {
	t = template.Must(t.Parse(overrideItemList))
	err := t.Execute(os.Stdout, items)
	if err != nil {
		log.Println("executing template:", err)
	}
}

func main() {
	fmt.Println("before override:")
	tmplBeforeOverride()
	fmt.Println("after override:")
	tmplafterOverride()
}
