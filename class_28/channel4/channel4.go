package main

import "fmt"

func main() {
	myChannel := make(chan string)
	go greetings(myChannel)
	//fmt.Println(<- myChannel)

	receiveValue := <- myChannel
	fmt.Println(receiveValue)
}

func greetings(myChannel chan string) {
	myChannel <- "Hi"
}