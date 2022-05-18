package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// func viewHandler(writer http.ResponseWriter, request *http.Request) {
// 	placeholder := []byte("signature list goes here")
// 	_, err := writer.Write(placeholder)
// 	check(err)
// }
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")
	fmt.Printf("%#v\n",signatures)
	html, err := template.ParseFiles("view.html")
	check(err)
	err = html.Execute(writer, nil)
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