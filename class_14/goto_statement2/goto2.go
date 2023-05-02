package main

import (
	"fmt"
)

func main(){
	fmt.Println("Before execution of goto statement.")
	i:=1
	START: 	
		for i<=5 {
			if i%2 == 0 {
				i++
				goto START //we can achieve the same goal using "continue statement"
				// continue
			}
			fmt.Println(i)
			i++
		}
	fmt.Println("after execution.")
}

