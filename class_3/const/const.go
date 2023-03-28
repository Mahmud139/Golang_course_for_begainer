package main

import "fmt"

const(
	year = 365
	leapYear int32 = 366
	//leapYear = int32(366)
)

// firstname 
// var firstNameMahmud string
// var first_name string
// var name12 string



func main() {
	hours := 24
	//minutes := int32(60)
	var minutes int32 = 60
	fmt.Println(hours * year) /* In this case, hours is an int, and years is untyped. 
	When the program compiles, it explicitly converts years to an int, which allows 
	the multiplication operation to succeed. */
	fmt.Println(year * minutes) /* In this case, minutes is an int32, and year is untyped. 
	When the program compiles, it explicitly converts years to an int32, which allows 
	the multiplication operation to succeed. */
	fmt.Println(minutes * leapYear) /* In this case, minutes is an int32, and leapYear 
	is a typed constant of int32. There is nothing for the compiler to do this time as
	both variables are already of the same type. */

	/* If we try to multiply two types that are typed and not compatible, 
	the program will not compile: */
	//fmt.Println(hours * leapYear) 
	/* In this case, hours was inferred as an int, and leapYear was explicitly declared 
	as an int32. Because Go is a typed language, an int and an int32 are not compatible 
	for mathematical operations. */

}

/* Here when we declare the constant leapYear, we define it as data type int32. 
Therefore it is a typed constant, which means it can only operate with int32 data types. 
The year constant we declare with no type, so it is considered untyped. Because of this, 
you can use it with any integer data type. */