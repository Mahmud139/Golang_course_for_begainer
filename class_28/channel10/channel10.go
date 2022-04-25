//Updating our channel to carry a struct
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type page struct {
	url string
	size int
}

func responseSize(url string, myChannel chan page) {
	//fmt.Println("Getting",url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	myChannel <- page{url,len(body)}
}

func main() {
	pages := make(chan page)
	urls := []string{"https://example.com","https://golang.org/","https://golang.org/doc"}
	for _, url := range urls {
		go responseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s : %d\n",page.url,page.size)
	} 
}