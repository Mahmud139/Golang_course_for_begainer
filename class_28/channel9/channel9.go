//Fixing our web page size program with channels
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func responseSize(url string, myChannel chan int) {
	fmt.Println("Getting",url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	myChannel <- len(body)
}

func main() {
	size := make(chan int)
	go responseSize("https://example.com", size)
	go responseSize("https://golang.org/", size)
	go responseSize("https://golang.org/doc", size)
	fmt.Println(<- size)
	fmt.Println(<- size)
	fmt.Println(<- size) 
}