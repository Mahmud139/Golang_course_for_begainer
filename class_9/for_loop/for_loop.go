/*suppose we want to write a program that counts 1 to 10, starting from 1, with each
number on its own line.*/
package main

import "fmt"

func main(){
	// fmt.Println(1)
	// fmt.Println(2)
	// fmt.Println(3)
	// fmt.Println(4)
	// fmt.Println(5)
	// fmt.Println(6)
	// fmt.Println(7)
	// fmt.Println(8)
	// fmt.Println(9)
	// fmt.Println(10)

	// or like this way: 

// 	fmt.Println(`1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// 10`)

	/*But both of these programs are pretty tedious to write.
What we need is a way of doing something multiple times.*/

	/*Rewriting our previous program using a for statement looks like this:*/
	for i:= 1; i<=5; i++{
		fmt.Println(i)
	}

}

