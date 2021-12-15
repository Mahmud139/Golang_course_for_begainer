package main

import "fmt"

func main(){
	fmt.Println()
	slice := []int{12,12,12,12,12,12,12}
	result := average(slice...)
	fmt.Println(result)
}

func average(numbers ...int) int {
	sum := 0
	for _, number := range numbers{
		sum += number
	}
	return sum/len(numbers)
}