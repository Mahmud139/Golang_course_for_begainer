package main

import "fmt"

func main(){
	recursion()
}

func recursion(){
	fmt.Println("Oh, no, I'am stuck.")
	recursion()
}

//these types of recursion are called "infinite recursion"