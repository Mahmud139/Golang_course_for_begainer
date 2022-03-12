package main

import "fmt"

type part struct {
	description string 
	count int
}

func main(){
	p := minimumOrder("Hex bolts")
	fmt.Println(p.description, p.count)
}

func minimumOrder(description string) part {
	var p part 
	p.description = description
	p.count = 12
	return p
}