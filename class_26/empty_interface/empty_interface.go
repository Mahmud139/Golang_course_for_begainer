package main

import "fmt"

// type Anything interface {

// }

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
	//thing.MakeSound()
	horn, ok := thing.(Horn)
	if ok {
		horn.MakeSound()
	}
}

type Horn string
func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

func AcceptHorn(horn Horn) {
	fmt.Println(horn)
	horn.MakeSound()
}

func main(){
	//fmt.Println(12, 12.121, "Mahmud", true)

	AcceptAnything(321.12121)
	AcceptAnything("Hasan")
	AcceptAnything(Horn("Hasan"))
}