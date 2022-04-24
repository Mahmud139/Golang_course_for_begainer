//Observing goroutine synchronization
package main

import (
	"fmt"
)

func send(myChannel chan string) {
	fmt.Println("***sending value***")
	myChannel <- "a"
	fmt.Println("***sending value***")
	myChannel <- "b"
}

func main() {
	myChannel := make(chan string)
	go send(myChannel)
	fmt.Println(<- myChannel)
	fmt.Println(<- myChannel)
}