package main

import "fmt"

type part struct {
	description string
	count int
}

type car struct {
	name string
	topSpeed float64
}

func main(){
	var porsche car 
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println("name: ", porsche.name)
	fmt.Println("topspeed: ", porsche.topSpeed)

	var bolt part
	bolt.description = "Hex bolts"
	bolt.count = 24
	fmt.Println("Description:", bolt.description)
	fmt.Println("Count: ", bolt.count)
}