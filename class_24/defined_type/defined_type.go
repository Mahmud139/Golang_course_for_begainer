package main

import "fmt"

type Liters float64
type Gallons float64

func main(){
	var carFuel Gallons
	var busFuel Liters
	// carFuel = 12.12
	// busFuel = 34.23
	// fmt.Printf("%T\n",carFuel)
	// fmt.Printf("%T",busFuel)
	carFuel = Gallons(10.0)
	busFuel = Liters(240.0)
	fmt.Printf("%T\n",carFuel)
	fmt.Printf("%T\n",busFuel)
	fmt.Println(carFuel, busFuel)

	//Use shorthand variable declaration
	// carFuel := Gallons(10.0)
	// busFuel := Liters(240.0)

	// carFuel = Liters(10.1)
	// busFuel = Gallons(240.2)
	/*If you have a variable that uses a defined type, you cannot assign a value of a
	different defined type to it, even if the other type has the same underlying type.
	This helps protect developers from confusing the two types.*/


	/*But you can convert between types that have the same underlying type. So
	Liters can be converted to Gallons and vice versa, because both have an
	underlying type of float64.*/
	carFuel = Gallons(Liters(240.0))
	busFuel = Liters(Gallons(10.0))

}