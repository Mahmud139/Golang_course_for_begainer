package main

import (
	"fmt"
	//"reflect"
	"strconv"
)

func main(){
	// var myint float64 = 2.12
	// var n int = 9
	// conv := int(myint) // converting float to int
	// conv2 := float64(n) // converting int to float
	// conv3 := int16(n)  // converting int to int16
	

	// fmt.Println(reflect.TypeOf(conv))
	// fmt.Println(reflect.TypeOf(conv2))
	// fmt.Println(reflect.TypeOf(conv3))

	//converting numbers to strings
	// var num string = "121212"
	// res:= "shihab" + num
	//var n int = 6512
	// var conv4 = rune(a)
	// fmt.Println(reflect.TypeOf(res))
	// fmt.Println(res)
	
	//a:= fmt.Sprint(n)
	//fmt.Println(reflect.TypeOf(a))
	//fmt.Println(a)

	// a:= strconv.Itoa(12)
	// fmt.Println(reflect.TypeOf(a))
	// fmt.Println(a)


	// converting string to numbers

	//var s string = "123"
	var f string = "123.121"
	// res, err:= strconv.Atoi(s) // converting string to integer number
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(reflect.TypeOf(res))
	// fmt.Println(res)

	res2, err := strconv.ParseFloat(f,32) // converting string to float32 number
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res2)

	// Now try to convert a string that is not a number:





	// why this conversion topic is soo important?

	// user := "shihab"
	// lines := 50
	// fmt.Println("Congratulations, " + user + "! You just wrote " + lines + " lines of code.") 
}