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
    //             break
    //         }
    //     }
    // }


	//solution 
	outer:  
		for i := 0; i < 3; i++ {
				for j := 1; j < 4; j++ {
					fmt.Printf("i = %d , j = %d\n", i, j)
					if j == 1 {
						break outer
					}
				}
		}
}