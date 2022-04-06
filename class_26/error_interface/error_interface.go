package main

import "fmt"

type CommedyError string
func (c CommedyError) Error() string {
	return string(c)
}

func main() {
	// err := fmt.Errorf("a height of %0.2f is invalid", -2.333)
	// fmt.Println(err.Error())
	// fmt.Println(err)

	var err error 
	err = CommedyError("What's a programmer's favorite food?")
	fmt.Println(err)
}