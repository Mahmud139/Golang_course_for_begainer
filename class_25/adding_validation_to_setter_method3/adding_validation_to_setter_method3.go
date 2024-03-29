package main

import (
	"fmt"
	"golang_tutorial/class_25/calendar"
	"log"
)

func main(){
	date := calendar.Date{}
	//err := date.SetYear(0)
	err := date.SetYear(2019)
	checkError(err)
	err = date.SetMonth(6)
	checkError(err)
	err = date.SetDay(6)
	checkError(err)
	/*Unexported variables, struct fields, functions, and methods can still be
	accessed by exported functions and methods in the same package.*/
	fmt.Println(date)

	/*if we try to update main.go to print an
individual Date field, we won’t be able to access it!*/
	//fmt.Println(date.year)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}