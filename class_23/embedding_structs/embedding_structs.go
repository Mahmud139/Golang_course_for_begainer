package main

import (
	"fmt"
	"golang_tutorial/excercise/magazine2"
)

func main(){
	subscriber := magazine2.Subscriber{Name: "mahmud"}
	subscriber.Street = "1242"
	subscriber.State = "narsingdi"
	subscriber.City = "miami"
	subscriber.PostalCode = "121"
	fmt.Println(subscriber)

	employee := magazine2.Employee{Name: "hasan"}
	employee.City = "dhaka"
	employee.PostalCode = "232"
	employee.State = "MA"
	employee.Street = "3432"
	fmt.Println(employee)
}