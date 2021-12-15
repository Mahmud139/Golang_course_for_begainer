//find the gretest value 
package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Println(maximum(71.8,90.3,21.5))
	fmt.Println(maximum(12.34, 43.45,99.99,23.23,56.43,78.90))
}

func maximum(numbers ...float64) float64{
	max := math.Inf(-1)
	for _, number := range numbers{
		if number > max {
			max = number
		}
	}
	return max
}