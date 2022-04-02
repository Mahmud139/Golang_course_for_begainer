package main

import "fmt"

type Whistel string

func (w Whistel) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type NoiseMaker interface {
	MakeSound()
}

func play(n NoiseMaker) {
	n.MakeSound()
}

func main() {
	play(Whistel("mahmud"))
	play(Horn("hasan"))
}