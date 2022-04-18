package main

import "fmt"

func main() {
	one()
}

func one() {
	defer fmt.Println("defered in one()")
	two()
}

func two(){
	defer fmt.Println("defered in two()")
	panic("Let's see what's been defered!")
}

