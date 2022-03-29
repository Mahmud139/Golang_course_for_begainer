package main

import (
	"fmt"
	"golang_tutorial/class_25/calendar"
)

func main(){
	event := calendar.Event{}
	fmt.Println(event.Year())
	//unexported fields don't get promoted to outer struct type
	// event.month = 5
	// event.Date.year = 10
}