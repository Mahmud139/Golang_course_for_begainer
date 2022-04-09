package main

import "fmt"

// type Anything interface {

// }

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
}

func main(){
	//fmt.Println(12, 12.121, "Mahmud", true)

	AcceptAnything(321.12121)
	AcceptAnything("Hasan")
}