package main

import "fmt"

func Socialize() {
	defer fmt.Println("Goodbye")
	fmt.Println("Hello!")
	fmt.Println("Nice Weather, eh?")
}

func main(){
	Socialize()
}