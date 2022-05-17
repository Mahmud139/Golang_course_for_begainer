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

func main() {
	templateText := "Template start\nAction: {{.}}\nTemplate end\n"
	tmpl, err := template.New("test").Parse(templateText)
	check(err)
	err = tmpl.Execute(os.Stdout, "ABC")
	check(err)
	err = tmpl.Execute(os.Stdout, 1234)
	check(err)
	err = tmpl.Execute(os.Stdout, true)
	check(err)
}