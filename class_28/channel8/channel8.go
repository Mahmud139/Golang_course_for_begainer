//Observing goroutine synchronization
package main

import (
	"fmt"
	// "time"
)

func send(myChannel chan string) {
	fmt.Println("***sending value***")
	myChannel <- "a"
	fmt.Println("***sending value***")
	myChannel <- "b"
	fmt.Println("***Finish!***")
	/*this won't run because of as soon as main function receives both channel value it done executing, So there's no time or opportunity for this line of code to execute.*/
}

func main() {
	myChannel := make(chan string)
	go send(myChannel)
	fmt.Println(<- myChannel)
	fmt.Println(<- myChannel)
	// time.Sleep(1 * time.Second)
	fmt.Println("received!")
}