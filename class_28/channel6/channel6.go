//Synchronizing goroutines with channels
package main

import (
	"fmt"
)

func main() {
	myChannel1 := make(chan string)
	myChannel2 := make(chan string)
	go abc(myChannel1)
	go def(myChannel2)
	fmt.Print(<- myChannel1)
	fmt.Print(<- myChannel2)
	fmt.Print(<- myChannel1)
	fmt.Print(<- myChannel2)
	fmt.Print(<- myChannel1)
	fmt.Print(<- myChannel2)
}

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}