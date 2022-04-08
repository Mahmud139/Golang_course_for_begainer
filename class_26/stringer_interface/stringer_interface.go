package main

import "fmt"

type Gallons float64
type Liters float64
type Milliliters float64

type Coffepot string
func (c Coffepot) String() string {
	return string(c) + " coffe pot"
}

func main(){
	// fmt.Println(Gallons(12.12121))
	// fmt.Println(Liters(12.12121))
	// fmt.Println(Milliliters(12.12121))

	// fmt.Printf("%0.2f gal\n",Gallons(12.1212))
	// fmt.Printf("%0.2f L\n",Liters(12.12121))
	// fmt.Printf("%0.2f ml\n",Milliliters(12.12121))

	coffepot := Coffepot("Mahmud")
	fmt.Println(coffepot.String())
	fmt.Println(coffepot)
	fmt.Print(coffepot, "\n")
	fmt.Printf("%s \n",coffepot)
}