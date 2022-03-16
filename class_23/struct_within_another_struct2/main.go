package main

import (
	"fmt"
	"golang_tutorial/excercise/magazine"
)

func main(){
	subscriber := magazine.Subscriber{}
	fmt.Printf("%#v",subscriber.HomeAddress)

	subscriber.HomeAddress.State = "FA"
	fmt.Println(subscriber.HomeAddress.State)
}