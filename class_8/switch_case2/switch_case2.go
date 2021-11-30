//Both switch statement and switch expression are present
package main

import "fmt"

func main(){
	
	switch ch:= "a"; ch {  //we have switch statement having a short declaration.
	case "a":
		fmt.Println("a")
	case "b":
		fmt.Println("b")
	case "c", "d":
		fmt.Println("c or d")
	default:
		fmt.Println("No matching charecter.")
	}


	//fmt.Println(ch)
}

//ch variable is only available within switch block.