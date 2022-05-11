package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", viewHandler)
	err := http.ListenAndServe("localhost:8080",nil)
	if err != nil {
		log.Fatal(err)
	}
	//http.HandleFunc("/hello", viewHandler) //if we write here the handlefunc it won't work
	// because when we start our server it's just start listening on that stage, the statement
	// below couldn't be ran on that time. 
}

func viewHandler(writter http.ResponseWriter, request *http.Request) {
	message := []byte("Hello, web!")
	_, err := writter.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}