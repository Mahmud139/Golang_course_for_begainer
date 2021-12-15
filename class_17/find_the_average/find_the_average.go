package main

import "fmt"

func main(){
	fmt.Println(average(12,12,12,12,12,12,12))
	fmt.Println(average(1,2,3,4,5,6,7,8,9,10))
}

func average(numbers ...int) int {
	sum := 0
	for _, number := range numbers{
		sum += number
	}
	return sum/len(numbers)
}