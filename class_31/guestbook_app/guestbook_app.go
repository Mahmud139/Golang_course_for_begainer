package main

import (
	"log"
	"net/http"
	"html/template"
)

// func viewHandler(writer http.ResponseWriter, request *http.Request) {
// 	placeholder := []byte("signature list goes here")
// 	_, err := writer.Write(placeholder)
// 	check(err)
// }
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("view.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	check(err)
}