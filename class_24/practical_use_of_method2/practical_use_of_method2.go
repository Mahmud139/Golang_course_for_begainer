package main

import "fmt"

type Liters float64
type Milliliter float64
type Gallons float64

func main() {
	milk := Gallons(2)
	fmt.Printf("%0.2f gallons equals %0.3f Liter.\n",milk, milk.ToLiters())
	fmt.Printf("%0.2f gallons equals %0.3f Milliliter.\n",milk, milk.ToMilliliters())
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliliters() Milliliter {
	return Milliliter(g * 3785.41)
}