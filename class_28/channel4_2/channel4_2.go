package main

import "fmt"

func main() { //this main goroutine is our receiving goroutine
	myChannel := make(chan string)
	go greetings(myChannel)
	receiveValue := <- myChannel

	fmt.Println("This line won't print out until receiving goroutine receive the data")
	fmt.Println(receiveValue)
}

func greetings(myChannel chan string) { //this is our sending goroutine
	// myChannel <- "Hi"
}

/*A receive operation blocks the receiving goroutine until another goroutine executes a send operation on the same channel. */