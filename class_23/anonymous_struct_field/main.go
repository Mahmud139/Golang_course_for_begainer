package main

import (
	"fmt"
	"golang_tutorial/excercise/magazine2"
)

func main(){
	subscriber := magazine2.Subscriber{Name: "mahmud"}
	subscriber.Address.City = "dhaka"
	subscriber.Address.PostalCode = "12323.1231"
	fmt.Println(subscriber.Address.City)
	fmt.Println(subscriber.Address.PostalCode)
}