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

	date = Date{Year: 2019, Month: 14, Day: 50}
	fmt.Println(date)
	date = Date{Year: 0, Month: 0, Day: -2}
	fmt.Println(date)
	date = Date{Year: -999, Month: -1, Day: 0}
	fmt.Println(date)
}