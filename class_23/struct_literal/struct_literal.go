package main

import (
	"fmt"
	"golang_tutorial/excercise/magazine"
)

func main(){
	//here i have wrote 3 types of struct literals writing method
	subscriber := magazine.Subscriber{Name: "mahmudul Hasan", Rate: 3.00, Active: true}

	// subscriber := magazine.Subscriber{
	// 	Name: "mahmudul Hasan",
	// 	Rate: 3.00,
	// 	Active: true,
	// }

	// subscriber := magazine.Subscriber{
	// 	Name: "mahmudul Hasan",
	// 	Rate: 3.00,
	// 	Active: true}

	fmt.Println(subscriber.Active)
	fmt.Println(subscriber.Name)
}