package main

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("view.html")
	check(err)
	err = html.Execute(w, nil)
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	log.Println("Server is listening on localhost:8080...")
	err := http.ListenAndServe("localhost:8080", nil)
	check(err)
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}