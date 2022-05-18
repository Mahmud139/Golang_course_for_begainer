//Repeating parts of a template with “range” actions
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
	text := ""
}

/*A section of a template between a {{range}} action and its 
corresponding {{end}} marker will be repeated for each value collected in an array, 
slice, map, or channel. Any actions within that section will also be repeated.

Within the repeated section, the value of dot will be set to the current element from 
the collection, allowing you to include each element in the output or do other 
processing with it.*/