package main

import "fmt"

func main() { //this main function is our receiving goroutine
	myChannel := make(chan string)
	go greetings(myChannel)
	receiveValue := <- myChannel

	fmt.Println("This line won't print out until receiving goroutine receive the data")
	fmt.Println(receiveValue)
}

func greetings(myChannel chan string) { //this is our sending goroutine
	// myChannel <- "Hi"
	fmt.Println("From greetings goroutine...")
}

/*A receive operation blocks the receiving goroutine until another goroutine executes a send operation on the same channel. */

/*Q: why we got dead lock error?
Ans: As we know that A receive operation blocks the receiving goroutine until another goroutine executes a send operation on the same channel. So in this situation our receiving goroutine is our main function, so the receive operation block this main function, when main function is blocked then there's nothing to execute. That's why we got deadlock error.*/