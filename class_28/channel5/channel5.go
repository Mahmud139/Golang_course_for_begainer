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

	fmt.Println(<- myChannel1)
	fmt.Println(<- myChannel2)
	fmt.Println(<- myChannel1)
	fmt.Println(<- myChannel2)
	fmt.Println(<- myChannel1)
	fmt.Println(<- myChannel2)
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

/*We know what the order will be because the abc goroutine blocks each time it sends a value to a channel until the main goroutine receives from it. The def goroutine does the same. The main goroutine becomes the orchestrator(rearranger) of the abc and def goroutines, allowing them to proceed only when it’s ready to read the values they’re sending.
*/
/*channels also ensure the sending goroutine has sent the value before the receiving channel attempts to use it.*/