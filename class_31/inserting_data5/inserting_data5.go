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

type Subscriber struct{
	Name string
	Rate float64
	Active bool
}

func main() {
	text := "Name: {{.Name}}\n{{if .Active}}Rate: ${{.Rate}}\n{{end}}"
	subscriber := Subscriber{Name: "Mahmud", Rate: 5.99, Active: true}
	executeTemplate(text, subscriber)
	subscriber = Subscriber{Name: "Mahmud", Rate: 5.99, Active: false}
	executeTemplate(text, subscriber)
}