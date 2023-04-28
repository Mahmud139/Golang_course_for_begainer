package main

import (  
    "fmt"
)

func main() {  
	//first part 
    // for i := 0; i < 3; i++ {
    //     for j := 1; j < 3; j++ {
    //         fmt.Printf("i = %d , j = %d\n", i, j)
    //         if i == j {
    //             break //here we want to break out the outer loop, but it's just 
						//break the current loop. what's the solution?
    //         }
    //     }
    // }


	//solution : we can use level to accomplish our requirement
	outer:  
		for i := 0; i < 3; i++ {
				for j := 1; j < 4; j++ {
					fmt.Printf("i = %d , j = %d\n", i, j)
					if i == j {
						break outer
					}
				}
		}
}