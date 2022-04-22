package main

import "fmt"

func main() {
	myChannel := make(chan string)
	go greetings(myChannel)
	fmt.Println(<- myChannel)
}

func greetings(myChannel chan string) {
	myChannel <- "Hi"
}