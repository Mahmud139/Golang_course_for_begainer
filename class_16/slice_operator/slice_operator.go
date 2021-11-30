package main

import (
	"fmt"
 	// "reflect"
)

func main(){
	//slide no 11
	// myArray := [6]int{1,2,3,4,5,6}
	// mySlice := myArray[1:3]
	// fmt.Println(mySlice)
	// fmt.Println(reflect.TypeOf(mySlice).Kind())

	//slide no 12
	// underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	// slice1 := underlyingArray[0:3]
	// fmt.Println(slice1)
	// fmt.Println(reflect.TypeOf(underlyingArray).Kind())
	// fmt.Println(reflect.TypeOf(slice1).Kind())
	// fmt.Printf("%T",slice1)


	// underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	// slice2 := underlyingArray[:4]
	// fmt.Println(slice2)

	// slide no 13
	// underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	// slice3 := underlyingArray[1:]
	// fmt.Println(slice3)

	// extra 
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	slice4 := underlyingArray[:]
	fmt.Println(slice4)
	fmt.Printf("%T",slice4)

}