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

/*When a program panics, all deferred function calls will still be made. If there’s
more than one deferred call, they’ll be made in the reverse of the order they were
deferred in.*/