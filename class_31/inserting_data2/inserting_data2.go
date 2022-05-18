//Making parts of a template optional with “if” actions
package main

import (
	"html/template"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func main() {
	text := "Start {{if .}}Dot is true!{{end}} Finish\n"
	executeTemplate(text, true)
	executeTemplate(text, false)
}

/*A section of a template between an {{if}} action and its 
corresponding {{end}} marker will be included only if a condition is true.*/