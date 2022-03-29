package main

import (
	"fmt"
	"golang_tutorial/class_25/calendar"
	"log"
)

func main(){
	event := calendar.Event{}
	// err := event.SetTitle("An extremely long and impractical title")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	err := event.SetTitle("My birthday")
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetYear(1999)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(31)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
}