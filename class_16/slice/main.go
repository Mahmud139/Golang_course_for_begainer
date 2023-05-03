package main

import (
	"fmt"
	// "reflect"
)

func main(){
	// var MyArray [3]string
	// var MySlice []string
	// fmt.Println(reflect.TypeOf(MyArray).Kind())
	// fmt.Println(reflect.TypeOf(MySlice).Kind())
	
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
	// s := []int{}//here we can't do indexing because we declare a 0 length slice
	//s[1] = 12 //here this statement is not valid, but append() is valid here
	// fmt.Println(s)
	// fmt.Println(len(s))

	// n := []int{} // eta empty thaka obostai amra indexing kore kiso assign korte parbo na but can use append()
	n := []int{1,2} // now we can assign at index 0,1 position. also can use append()
	fmt.Println(n)
	fmt.Println(cap(n))
	fmt.Println(len(n))
	// for i := 0; i < 2; i++ {
	// 	n = append(n, i)
	// 	//n[i] = i
	// } 

	// n := make([]int, 4)
	// fmt.Println(n)
	// fmt.Println(cap(n))
	// fmt.Println(len(n))
	// for i := 0; i < cap(n); i++ {
	// 	n = append(n, i) // eta likha jabe na.. karon error pabo amra...kno error pabo ta gobesona koro :)
	// 	//n[i] = i 
	// } 
	// fmt.Println(n)
	// fmt.Println(cap(n))
	// fmt.Println(len(n))

	// a := []int{1,2,3,4,5}
	// s := make([]int, 3, 3)
	// s = a[1:]
	// fmt.Print(s)

	//num := make([]int, 4)
	// num := []int{}
	// fmt.Println(len(num))
	// fmt.Println(cap(num))
	// for i := 0; i < 4; i++ {
	// 	//num[i] = i
	// 	num = append(num, i) // can we use this statement instead of using above statement?
	// }
	// fmt.Println(len(num))
	// fmt.Println(cap(num))
	// fmt.Println(num)
}