package main

import "fmt"

type Whistel string

// func (w Whistel) MakeSound() {
// 	fmt.Println("Tweet!")
// }

func (w Whistel) MakeSound() string {
	return fmt.Sprintf("Tweet! to %s",w)
}

//type Horn string

// func (h Horn) MakeSound() {
// 	fmt.Println("Honk!")
// }

type NoiseMaker interface {
	MakeSound() string
}

// func play(n NoiseMaker) {
// 	n.MakeSound()
// }

type test float64
func (t test) Error() string {
	return fmt.Sprintf("test %f form error", t)
}

func main() {
	var toy NoiseMaker = Whistel("mahmud")
	fmt.Println(toy)

	var err error = test(12.1)
	fmt.Println(err)
	//toy.MakeSound()

	// toy = Horn("hasan")
	// toy.MakeSound()

	// play(Whistel("mahmud"))
	// play(Horn("hasan"))
}