package main

import "fmt"

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) SetYear(year int) {
	d.Year = year
}

func (d *Date) SetMonth(month int) {
	d.Month = month
}

func (d *Date) SetDay(day int) {
	d.Day = day
}

func main() {
	date := Date{}
	date.SetYear(2019)
	date.SetMonth(5)
	date.SetDay(27)
	fmt.Println(date)
}