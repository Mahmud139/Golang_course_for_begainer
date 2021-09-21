/*Scan scans text read from standard input, "storing successive space-separated values 
into successive arguments. Newlines count as space." It returns the number of items 
successfully scanned. If that is less than the number of arguments, err will report why.*/

package main 

import "fmt"

func main(){
	var name string
	var age int

	// fmt.Println("please enter your name and age")
	// fmt.Scan(&name, &age)
	// fmt.Println(name, age)
	// fmt.Printf("%T %T",name, age)

	fmt.Println("Please enter your name")
	fmt.Scan(&name)

	fmt.Println("enter your age")
	fmt.Scan(&age)

	fmt.Println(name, age)

	// we can also see the return value
}