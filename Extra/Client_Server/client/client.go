package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/goLinuxCloud", nil)
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	body := response.Body.(http.bodyEOFSignal)
	fmt.Printf("Type of readCloser is %T\n", body)
	fmt.Printf("%#v\n", body)

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))
}