//Inserting struct fields into a template with actions
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

/*If the value in dot is a struct, then an action with dot followed by a field name 
will insert that field’s value in the template. Here we create a Part struct type, 
then set up a template that will output a Part value’s Name and Count fields:*/