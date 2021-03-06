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
}

func main() {
	// templateText := "Template start\nAction: {{.}}\nTemplate end\n"
	// tmpl, err := template.New("test").Parse(templateText)
	// check(err)
	// err = tmpl.Execute(os.Stdout, "ABC")
	// check(err)
	// err = tmpl.Execute(os.Stdout, 1234)
	// check(err)
	// err = tmpl.Execute(os.Stdout, true)
	// check(err)
	executeTemplate("Dot is: {{.}}!\n", "ABC")
	executeTemplate("Dot is: {{.}}!\n", 123)
}