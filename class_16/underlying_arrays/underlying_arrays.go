package main

import "fmt"

func main(){
	// slide no 15
	array1 := [5]string{"a", "b", "c", "d", "e"}
	slice1 := array1[0:3]
	fmt.Println(slice1)

	array2 := [5]string{"f", "g", "h", "i", "j"}
	slice2 := array2[2:5]
	fmt.Println(slice2)

	// slide no 16
	array3 := [5]string{"a", "b", "c", "d", "e"}
	slice3 := array3[0:3]
	slice4 := array3[2:5]
	fmt.Println(slice3, slice4)
}

