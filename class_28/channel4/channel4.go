package main

import "fmt"
//var myChannel = make(chan string)

func main() {
	myChannel := make(chan string)
	go greetings(myChannel)
	//fmt.Println(<- myChannel)

	receiveValue := <- myChannel
	fmt.Println(receiveValue)
	//test()
}

func greetings(myChannel chan string) {
	myChannel <- "Hi"
}

// func test() {
// 	//myChannel := make(chan string)
// 	go greetings(myChannel)
// 	fmt.Println(<- myChannel)
// }