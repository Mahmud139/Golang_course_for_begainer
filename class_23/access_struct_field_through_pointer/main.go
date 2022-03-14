//access_struct_field_through_pointer
package main

import "fmt"

type mystruct struct{
	myfield int
}

func main(){
	var value mystruct
	value.myfield = 3
	var pointer *mystruct = &value
	//fmt.Println(*pointer.myfield) 
	/*If you write *pointer.myField, Go thinks that myField must contain a pointer.
It doesn’t, though, so an error results. To get this to work, you need to wrap
*pointer in parentheses.*/
	fmt.Println((*pointer).myfield) 
	fmt.Println(pointer.myfield)
}