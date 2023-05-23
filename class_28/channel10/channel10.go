//Updating our channel to carry a struct
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL string
	Size int
}

func responseSize(url string, myChannel chan Page) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	myChannel <- Page{URL: url,Size: len(body)}
}

func main() {
	pages := make(chan Page)
	urls := []string{"https://example.com","https://golang.org/","https://golang.org/doc"}
	for _, url := range urls {
		go responseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		Page := <-pages
		fmt.Printf("%s : %d\n",Page.URL,Page.Size)
	} 
}