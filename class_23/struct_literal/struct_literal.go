package main

import (
	"fmt"
	"golang_tutorial/excercise/magazine"
)

func main(){
	subscriber := magazine.Subscriber{
		Name: "mahmudul Hasan",
		Rate: 3.00,
		Active: true,
	}
	fmt.Println(subscriber.Active)
	fmt.Println(subscriber.Name)
}