package main

import (  
    "fmt"
)

func main() {  
    // n := 3
    // for i := 0; i < n; i++ {
    //     for j := 0; j <= i; j++ {
    //         fmt.Print("*")
    //     }
    //     fmt.Println()
    // }


	for i := 0; i < 3; i++ {
        fmt.Printf("Outer loop iteration %d\n", i)
        for j := 0; j < 2; j++ {
            fmt.Printf("i= %d j=%d\n", i, j)
        }
	}
}