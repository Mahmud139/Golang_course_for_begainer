package main

import "fmt"

func main(){
	// String as a range in the for loop
    // for i, j:= range "XabCd" {
	// 	fmt.Printf("The index number of %v is %d\n", j, i) 
	// 	//fmt.Printf("The index number of %U is %d\n", j, i) // to print unicode we use %U verbs
	// }

	//The for loop iterate over each character of string.
	// for range "Hello" {
	// 	fmt.Println("Hello")
	// }

	n := [5]int{1,2,3,4}
	for range n {
		fmt.Println("Hello Shihab Uddin, Vhat khaiso?")
	}
}