package main

import (
	"fmt"
	"math"
)

func main(){
	 cans, reminder := floatParts(1.26)
	 fmt.Println(cans, reminder)
}

func floatParts(number float64) (integerPart int, fractionalPart float64) {
	wholeNumber := math.Floor(number)
	integerPart = int(wholeNumber)
	fractionalPart = number - wholeNumber
	return
}