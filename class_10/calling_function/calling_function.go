package main

import "fmt"

func main(){
	sayHi()
	fmt.Println("after runing sayhi function.")
	shihab()
}

func sayHi(){
	fmt.Println("Hello, How are you?")
}


func shihab(){
	sayHi()
	fmt.Println("Hello, How are you?")
}