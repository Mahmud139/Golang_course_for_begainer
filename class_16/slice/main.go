package main

import (
	"fmt"
	"reflect"
)

func main(){
	var MyArray [3]string
	var MySlice []string
	fmt.Println(reflect.TypeOf(MyArray).Kind())
	fmt.Println(reflect.TypeOf(MySlice).Kind())
	
	//Declaring a slice variable
	// var n []int //declaring a slice variable doesn't automatically create a slice 
	// n = make([]int, 7)//we need to use built-in "make()" to create a slice
	// n[3] = 12 //after create a slice you can assign and retrieve its elements same as array
	// n[1] = 23
	// n[4] = 122
	// fmt.Println(n[2])
	// fmt.Println(n[4])
	// fmt.Println(n)

	//Declaring a slice with using short-hand declaration
	// a := make([]int, 4)
	// a[2] = 23
	// a[3] = 34
	// fmt.Println(a[0])
	// fmt.Println(a[2])
	// fmt.Println(a)

	//Slice Literals
	// x := []int{1,2,3,4,5}//using slice literals we don't need to call make()
	// fmt.Println(x[1], x[4])
	// fmt.Println(x)


	//Note:
	// s := []int{}//here we can't do indexing because we declare a 0 length/empty slice
	//s[1] = 12 //here this statement is not valid, but append() is valid here
	// fmt.Println(s)
	// fmt.Println(len(s))


	//Important
	// a := []int{1,2,3,4,5}
	// s := make([]int, 3, 3)
	// s = a[1:]
	/*no matter how long length and capacity we set while defining slice, its always changeable. we can put any number of length's slice into that variable.*/
	// fmt.Print(s)

	// num := make([]int, 4)
	// num := []int{}
	// fmt.Println(len(num))
	// fmt.Println(cap(num))
	// for i := 0; i < 4; i++ {
	// 	//num[i] = i //this statement can be work with line no. 52 statement.
	// 	num = append(num, i) // this statement can be work with line no. 52,53 statement
	// }
	// fmt.Println(len(num))
	// fmt.Println(cap(num))
	// fmt.Println(num)
}