package main

import "fmt"

func main(){
	//var name string
	// var myFunc func(string)
	// myFunc = sayHi
	//myFunc()

	//name := "Shihab"
	myFunc := sayHi
	myFunc()
	//a := sayHi()

}

func sayHi() {
	fmt.Println("Hello from sayHi function.")
}