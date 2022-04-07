package main

import (
	"fmt"
)

type CommedyError string
func (c CommedyError) Error() string {
	return string(c)
}

func main() {
	var err error = CommedyError("What's a programmer's favorite food?")
	fmt.Println(err)
}