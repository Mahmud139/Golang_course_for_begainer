package main

import (
	"fmt"
	"golang_tutorial/class_25/calendar"
	"log"
)

func main(){
	event := calendar.Event{}
	//unexported fields don't get promoted to outer struct type
	// event.month = 5
	// event.Date.year = 10
	err := event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(21)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
	
}