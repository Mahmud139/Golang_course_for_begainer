package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	text := "Here's my Template!\n"
	tmpl, err := template.New("text").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, nil)
	check(err)

	_, err = os.Stdout.Write([]byte("hello"))
	check(err)
}