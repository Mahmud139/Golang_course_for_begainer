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
	
	
	//var n []int
	// n := make([]int, 2, 4)
	// n = []int{1,2}
	//n := [5]int{1,2,1,1}
	//n[3] = 12
	// fmt.Println(n)
	// fmt.Println(cap(n))
	// fmt.Println(len(n))
	//n := make([]int, 4, 10)
	//n := []int{} // eta empty thaka obostai amra indexing kore kiso assign korte parbo na but can use append()
	//n := []int{1,2} // now we can assign at index 0,1 position. also can use append()
	// fmt.Println(n)
	// fmt.Println(cap(n))
	// fmt.Println(len(n))
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