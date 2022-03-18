package main

import "fmt"

type Liters float64
type Gallons float64

func main(){
	/*A defined type supports all the same operations as its underlying type. Types
	based on float64, for example, support arithmetic operators like +, -, *, and /,
	as well as comparison operators like ==, >, and <.*/
	
	fmt.Println(Liters(1.2) + Liters(3.4))
	fmt.Println(Gallons(5.5) - Gallons(2.2))
	fmt.Println(Liters(2.2) / Liters(1.1))
	fmt.Println(Liters(2.2) * Liters(1.2))
	fmt.Println(Gallons(1.1) == Gallons(1.1))
	fmt.Println(Liters(1.1) < Liters(3.3))
	fmt.Println(Liters(1.2) > Liters(3.4))

}