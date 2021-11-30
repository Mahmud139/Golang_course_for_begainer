package main

import "fmt"

func main() {
	arr1 := [4]int{0, 1, 2, 3}
	arr2 := [4]int{1, 2, 3, 4}
	arr3 := [4]int{1, 0, 1, 0}
	arr4 := [4]int{4, 3, 2, 1}
	arr5 := [4]int{0, 1, 2, 3}

	fmt.Println("arr1 == arr2  ", arr1 == arr2)
	fmt.Println("arr2 == arr3  ", arr2 == arr3)
	fmt.Println("arr3 == arr4  ", arr3 == arr4)
	fmt.Println("arr4 == arr1  ", arr4 == arr1)
	fmt.Println("arr1 == arr5  ", arr1 == arr5)
}