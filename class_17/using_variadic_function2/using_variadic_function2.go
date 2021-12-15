package main

import "fmt"

func main(){
	fmt.Println(inRange(1,100,12,23,1,45,6,76,56,89,90))
	fmt.Println(inRange(-10,10,-3,-1,23,45,567,67,-7,-9))
}

func inRange(min int, max int, numbers ...int) []int {
	var result []int
	for _, number := range numbers{
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}