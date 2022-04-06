package main

import "fmt"

func main() {
	err := fmt.Errorf("a height of %0.2f is invalid", -2.333)
	fmt.Println(err.Error())
	fmt.Println(err)
}