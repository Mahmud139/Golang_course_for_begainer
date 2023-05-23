package main

import "fmt"

func main() {
	myChannel := make(chan string)
	go greetings(myChannel)
	// receiveValue := <- myChannel
	// fmt.Println(receiveValue)
}

func greetings(myChannel chan string) { //this is our sending goroutine
	myChannel <- "Hi"
	fmt.Println("This line won't print out until receiving goroutine receive the data")
}

/*A send operation blocks the sending goroutine until another goroutine executes a receive operation on the same channel. */