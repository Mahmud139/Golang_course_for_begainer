package main

import "fmt"

type Gallons float64
func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal",g)
}

type Liters float64
func (l Liters) String() string {
	return fmt.Sprintf("%0.2f l",l)
}

type Milliliters float64
func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f ml",m)
}


func main() {
	fmt.Println(Gallons(12.1212))
	fmt.Println(Liters(12.1212))
	fmt.Println(Milliliters(12.12121))
}