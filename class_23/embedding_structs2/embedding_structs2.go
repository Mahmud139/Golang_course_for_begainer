package main

import (
	"fmt"
)

type Subscriber struct {
	Name string
	Rate float64
	Active bool
	address
}

type address struct {
	street string
	City string
	State string
	PostalCode string
}

func main(){
	subscriber := Subscriber{Name: "mahmud"}
	subscriber.street = "1242"
	subscriber.State = "narsingdi"
	subscriber.City = "miami"
	subscriber.PostalCode = "121"
	fmt.Println(subscriber)
}