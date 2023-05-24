//Responding differently for different resource paths
package main

import (
	"log"
	"net/http"
)

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, web")
}
func banglaHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "kire, web")
}
func arabicHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Assalamo Alaikom, web")
}
func write(write http.ResponseWriter, s string) {
	message := []byte(s)
	_, err := write.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello",englishHandler)
	http.HandleFunc("/kire", banglaHandler)
	http.HandleFunc("/namaste",arabicHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
