//check if array's elements exits
package main

import (
	"fmt"
	"reflect"
)

func main(){
	name := [3]string{"Mahmud", "Hasan", "Shihab"}
	fmt.Println(itemExits(name, "Mahmud"))

}

func itemExits (arrayType [3]string, item string) bool {
	arr := reflect.ValueOf(arrayType)
	
	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
		// fmt.Println("Invalid type!")
		// return false 
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).String() == item {
			return true
		}
	}
	return false
}