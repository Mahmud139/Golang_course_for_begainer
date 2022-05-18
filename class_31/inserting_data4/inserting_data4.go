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

func executeTemplate(text string, data interface{}){
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
}

type Part struct{
	Name string
	Count int
}

func main() {
	text := "Name: {{.Name}}\nCount: {{.Count}}\n"
	executeTemplate(text, Part{Name: "Bolt", Count: 12})
	executeTemplate(text, Part{Name: "Cables", Count: 23})

}