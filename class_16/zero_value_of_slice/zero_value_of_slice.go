package main

import "fmt"

func main(){
	// slide no 25
	// floatSlice := make([]float64,10)
	// boolSlice := make([]bool, 10)
	// fmt.Println(floatSlice, boolSlice[5])

	var intSlice []int
	var stringSlice []string
	var emptySlice []int = []int{} //empty slice
	emptySlice2 := []int{} //initialized an empty slice with slice literal
	fmt.Println(intSlice, stringSlice, emptySlice)
	fmt.Printf("intSlice: %#v, stringSlice: %#v, emptySlice: %#v, emptySlice2: %#v\n",intSlice, stringSlice, emptySlice, emptySlice2)

	//slide no 26
	//intSlice[0] = 5
	//fmt.Println(intSlice)
	// fmt.Println(len(intSlice))

	// intSlice = append(intSlice, 25)
	// fmt.Printf("intSlice: %#v\n",intSlice)
	// fmt.Println(len(intSlice), cap(intSlice))


	//slide no 27
	// var slice []string
	// if len(slice) == 0 {
	// 	slice = append(slice, "first item")
	// }
	// fmt.Printf("%#v\n",slice)
}