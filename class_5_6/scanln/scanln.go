/*Scanln is similar to Scan, but "stops scanning at a newline and after
 the final item there must be a newline or EOF".*/
 // in golang EOF is The Enter key
package main 

import "fmt"

func main(){
	var name string
	var age int

	// fmt.Println("please enter your name and age")
	// fmt.Scanln(&name, &age)
	// fmt.Println(name, age)

	fmt.Println("Please enter your name")
	fmt.Scanln(&name)

	fmt.Println("enter your age")
	fmt.Scanln(&age)

	fmt.Println(name, age)

	// var input rune
	// _, _ = fmt.Scanf("%c", &input)
	// fmt.Printf("value: %v\n", input)

	// we can also see the return value
}