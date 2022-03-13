package main

import "fmt"

type subscriber struct {
	name string
	rate float64
	active bool
}

func main(){
	var s subscriber
	s.rate = 5.89
	applyDiscount(&s)
	fmt.Println(s.rate)
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}