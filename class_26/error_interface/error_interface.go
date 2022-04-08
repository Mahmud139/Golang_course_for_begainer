package main

import (
	"fmt"
	"log"
)

type OverheatError float64
func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
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
	// fmt.Println(err)

	var err error = CheckTemperature(121.45, 100.0)
	if err != nil {
		log.Fatal(err)
	}
}