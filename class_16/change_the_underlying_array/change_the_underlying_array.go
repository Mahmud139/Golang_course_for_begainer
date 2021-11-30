package main

import "fmt"

func main(){
	// slide no 17
	// array1 := [5]string{"a", "b", "c", "d", "e"}
	// slice1 := array1[0:3]
	// fmt.Println(slice1)
	// array1[1] = "X"
	// fmt.Println(array1)
	// fmt.Println(slice1)

	//slide no 18
	// array2 := [5]string{"f", "g", "h", "i", "j"}
	// slice2 := array2[2:5]
	// fmt.Println(slice2)
	// slice2[1] = "X"
	// fmt.Println(array2)
	// fmt.Println(slice2)


	array3 := [5]string{"a", "b", "c", "d", "e"}
	slice3 := array3[0:3]
	slice4 := array3[2:5]
	fmt.Println(slice3, slice4)
	slice3[2] = "X"
	//array3[2] = "X"
	fmt.Println(array3)
	fmt.Println(slice3, slice4)
}