package main

import "fmt"

type Number int

func (n Number) Double() {
	n *= 2
}

func main() {
	number := Number(12)
	fmt.Println("Original value of number:", number)
	number.Double()
	fmt.Println("number after calling double:", number)
}