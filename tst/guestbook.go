package main

import (
	"html/template"
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("view.html")
	check(err)
	_, _ = w.Write([]byte("helo"))
	err = html.Execute(w, nil)
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	log.Println("Server is listening on localhost:8080...")
	err := http.ListenAndServe("localhost:8080", nil)
	check(err)
}