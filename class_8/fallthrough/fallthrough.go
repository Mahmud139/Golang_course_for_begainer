package main

import "fmt"

func main(){
	i := 23
	// switch {
	// case i < 10 :
	// 	fmt.Println("i is less than 10.")
	// 	fallthrough
	// case i < 50 :
	// 	fmt.Println("i is less than 50.")
	// 	fallthrough
	// case i < 100 : 
	// 	fmt.Println("i is less than 100.")
	// 	fallthrough
	// default:
	// 	fmt.Println("i is greater than 100.")
	// }


/*fallthrough cannot be used in the last case of a switch since there are no more 
cases to fallthrough.*/

	switch {
	case i < 10 :
		fmt.Println("i is less than 10.")
		fallthrough
		//fmt.Println("Not allowed")
	case i < 50 :
		fmt.Println("i is less than 50.")
		fallthrough
	case i < 100 : 
		fmt.Println("i is less than 100.")
		fallthrough
	default:
		fmt.Println("i is greater than 100.")
	}
}

//Fallthrough will happen even when the case evaluates to false.

/*fallthrough should be the last statement in a case. If it is present somewhere 
in the middle, the compiler will complain that fallthrough statement out of place.*/



