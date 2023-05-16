package main

import (
	"fmt"
	"log"
)

type OverheatError float64
func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by mahmud %0.2f degrees!", o)
}

func CheckTemperature(actual float64, safe float64) error {
	excess := actual - safe 
	if excess > 0 {
		return OverheatError(excess)
	} 
	return nil 
}

func main() {
	// err := fmt.Errorf("a height of %0.2f is invalid", -2.333)
	// fmt.Println(err.Error())
	// fmt.Println(err) //here Go automatically call Error() method itself

	var err error = CheckTemperature(121.45, 100.0)
	// var err error = OverheatError(21.45)
	if err != nil {
		log.Fatal(err)
	}

	// var a int = 12
	// var b int = a
}

/*In Go, it is not mandatory to use the Error method specifically when printing errors. When an error is printed using the fmt.Println or fmt.Printf functions, Go automatically calls the Error method if the error value implements the error interface. */