package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	fmt.Println(reflect.TypeOf(&myInt))
	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat))
	var myBool bool 
	fmt.Println(reflect.TypeOf(&myBool))

	var pointer *int
	fmt.Println(reflect.TypeOf(pointer))
	fmt.Println(reflect.TypeOf(pointer).Kind())
}