package main

import (
	"fmt"
	//"reflect"
)

func main(){
	//declaring a single variable 
	/* 
	var name string 
	name = "I am shihab"
	fmt.Println(name) 
	*/
	
	// declaring a variable with initial value
	/*
	var name string = "I am shihab"
	fmt.Println(name)
	name = "I am mahmud"
	fmt.Println(name)
	*/

	// type inference 
	/*
	var name = "I am shihab"
	fmt.Println(name)
	fmt.Println(reflect.TypeOf(12))
	*/

	//multiple variable declaration
	// var n, h int 
	// n = 12 
	// h = 2323
	// fmt.Println(n,h)

	/*
	var (
		name string
		n int
		f float32
	)

	name = "shihab"
	n = 12
	f = 12.90
	fmt.Println(name, n, f)
    */

	//shorthand declaration
	/*
	name:= "shihab"
	fmt.Println(name)
	name = "shihab"
	fmt.Println(name)
	*/

	name, num := "shihab", 12
	fmt.Println(name, num)
	fmt.Println(len(name))
	name = "马姆"
	fmt.Println(len(name))
	/* here we use len function to count the "name" string.English characters needs 1 byte to store in memory. UTF8 charactes need 1-4 bytes. we know that if we want
	to use a function first we need to import the specific package. here we didn't do that.
	because In go, golang provides some predeclared identifiers. len function is one of them. 
	these predecleared identifiers are documented in builtin package: only for documentation
	purpose. that's why we don't need to import it. Is imported by default. */ 

	const x int = 12
	fmt.Println(x)
	// x = 13 // you cann't modify const variable once they have been declared.
	const y = 14 //untyped const variable, because we didn't specify the type of this variable
	/* This can be useful when working with numbers such as integer-type data. If the constant is untyped, it is explicitly converted, where typed constants are not. */
	fmt.Println(y)

	// String Concatenation
	// name:= "shihab "
	// name2:= "uddin"
	// res:= name + name2
	// fmt.Println(res)
}
