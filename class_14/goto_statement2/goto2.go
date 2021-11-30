package main

import (
	"fmt"

)

func main(){
	fmt.Println("Befor execution of goto statement.")
	i:=1
	START: 	
		for i<=5 {
			if i%2 == 0 {
				i++
				goto START
			}
			fmt.Println(i)
			i++
		}
	fmt.Println("after execution.")
}

