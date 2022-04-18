package main

import "fmt"

func main(){
	freckOut()
	fmt.Println("Exiting Normally")
}

func freckOut() {
	defer calmDown()
	panic("Oh no!")
	//fmt.Println("I won't be run")
}

func calmDown() {
	recover()
}