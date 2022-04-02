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

type Robot string 

func (r Robot) MakeSound() {
	fmt.Println("Beep Beep")
}

func (r Robot) Walk() {
	fmt.Println("Powering Legs")
}

type NoiseMaker interface {
	MakeSound()
}

func play(n NoiseMaker) {
	n.MakeSound() //ok, part of NoiseMaker Interface
	//n.Walk() //not ok, not part of NoiseMaker interface
}

func main() {
	// play(Whistel("mahmud"))
	// play(Horn("hasan"))

	play(Robot("Botoco Amber"))
}