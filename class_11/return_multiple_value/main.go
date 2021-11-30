package main

import "fmt"

func main(){
	//multiple return values without parameters
	///a, b, c := 1, false, "mahmud"
	myInt, myBool, myString := manyReturns()
	fmt.Println(myInt, myBool, myString)

	// multiple return values with parameters
	myInt, myBool, myString = manyReturns2(12, false, "Shihab's")
	fmt.Println(myInt, myBool, myString)
}

func manyReturns() (int, bool, string) {
	return 1, true, "Mahmud"
}

func manyReturns2(i int, b bool, s string) (int, bool, string) {
	return i+1, b && true, "Mahmud is "+s+ " friend"
}