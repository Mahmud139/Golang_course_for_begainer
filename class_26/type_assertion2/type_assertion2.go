package main

import "fmt"

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

func main(){
	var noisemaker NoiseMaker = Robot("hasan")
	noisemaker.MakeSound()
	var robot Robot = noisemaker.(Robot)
	robot.Walk()
	// noisemaker.(Robot).Walk()
}