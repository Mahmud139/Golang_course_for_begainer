//Only switch expression is present
package main

import "fmt"

func main(){
	ch:= "b"
	switch  ch {  
	case "a":
		fmt.Println("a")
	case "b":
		fmt.Println("b")
	case "c":
		fmt.Println("c or d")
	default:
		fmt.Println("No matching charecter.")
	}
}