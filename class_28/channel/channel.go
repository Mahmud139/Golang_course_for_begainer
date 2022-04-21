//Go statements can’t be used with return values
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) int {
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
	return len(body)
}

func main() {
	var size int
	//size = go responseSize("https://example.com")
	fmt.Println(size)
	//size = go responseSize("https://golang.org/")
	fmt.Println(size)
	//size = go responseSize("https://golang.org/doc")
	fmt.Println(size)
	time.Sleep(7*time.Second)
}