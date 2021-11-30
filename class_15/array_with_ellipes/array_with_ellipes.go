package main

import (
	"fmt"
	"reflect"
)

func main(){
	//x := [...]int{1,2,3,4,6,7}
	//var name string
	//var y [6]int

	// fmt.Println(len(x))
	// fmt.Println(reflect.TypeOf(x).Kind())
	// fmt.Printf("%T\n", x)
	//fmt.Printf("%T", name)

	//this is how we can ignore the length of the array in the declaration and replace it with "..."
	x := [...]int{12,12}
	//var x  = [...]int{12,12}
	//var y [3]int = [3]int{1,2,1}
	//var x [3]int
	
	fmt.Println(len(x))
	fmt.Println(reflect.TypeOf(x).Kind())
	fmt.Printf("%T\n", x)
}