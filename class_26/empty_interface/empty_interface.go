package main

import "fmt"

type Anything interface {}

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

	// var any Anything = Horn("poooooo")
	// any.MakeSound()
	/*Don’t try calling any methods on an empty-interface value! Remember, if you have a value with an interface type, you can only call methods on it that are part of the interface. And the empty interface doesn’t have any methods. That means there are no methods you can call on a value with the empty interface type!*/


	AcceptAnything(321.12121)
	AcceptAnything("Hasan")
	AcceptAnything(Horn("Hasan"))
}