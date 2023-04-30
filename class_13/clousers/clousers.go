package main

import "fmt"

func main() {
    modulus := getModulus()
	fmt.Println(modulus(-1))
    fmt.Println(modulus(2))
    fmt.Println(modulus(-5))

    
	test()
	test()
}

func getModulus() func( int) int {
    count := 0
    return func(x int) int {
        count = count + 1
        fmt.Printf("modulus function called %d times\n", count)
        if x < 0 {
            x = x * -1
        }
        return x
    }
}

//This function don't retain the current value between different function call
func test () {
	c := 0
	c = c + 1
	fmt.Println("count ", c)
}