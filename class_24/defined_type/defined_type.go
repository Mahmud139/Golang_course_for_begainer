package main

import "fmt"

type Liters float64
type Gallons float64

func main(){
	var carFuel Gallons
	var busFuel Liters
	carFuel = 12.12
	busFuel = 34.23
	fmt.Printf("%T\n",carFuel)
	fmt.Printf("%T",busFuel)

}