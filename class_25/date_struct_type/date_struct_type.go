package main

import "fmt"

type Date struct {
	Year  int
	Month int
	Day   int
}

func main() {
	date := Date{Year: 2019, Month: 12, Day: 23}
	fmt.Println(date)
}