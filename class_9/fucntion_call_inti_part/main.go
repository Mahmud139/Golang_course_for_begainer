package main

import "fmt"

func main() {
    i := 1
    //Function call in the init part in for loop
    for test(); i < 3; i++ {
        fmt.Println(i)
    }

    //Assignment in the init part in for loop
    for i = 2; i < 3; i++ {
        fmt.Println(i)
    }
}
func test() {
    fmt.Println("In test function")
}