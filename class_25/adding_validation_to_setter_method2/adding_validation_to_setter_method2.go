package main

import (
	"fmt"
	"golang_tutorial/class_25/calendar"
)
func main(){
	date := calendar.Date{}
	date.Year = 2019
	date.Month = 14
	date.Day = 50
	fmt.Println(date)

	date = calendar.Date{Year: 0, Month: 0, Day: -2}
	fmt.Println(date)
}