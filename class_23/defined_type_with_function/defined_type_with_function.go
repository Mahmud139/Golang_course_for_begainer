package main

import "fmt"

type part struct {
	description string 
	count int
}

func main(){
	var bolt part 
	bolt.description = "Hex bolts"
	bolt.count = 12

	showInfo(bolt)

}

func showInfo(p part) {
	fmt.Println("Description:",p.description)
	fmt.Println("Count:", p.count)
}