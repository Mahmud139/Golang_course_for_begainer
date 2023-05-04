package main

import "fmt"
    

func subtactOne(numbers []int) {  
    for i := range numbers {
        numbers[i] -= 2
    }

}
func main() {  
    nos := []int{8, 7, 6}
    fmt.Println("slice before function call", nos)
    subtactOne(nos)                               
    fmt.Println("slice after function call", nos) 
}

/*A slice contains the length, capacity and a pointer to the zeroth element of the underlying array. When a slice is passed to a function, even though it's passed by value, the pointer variable will refer to the same underlying array. Hence when a slice is passed to a function as parameter, changes made inside the function are visible outside the function too. Let's write a program to check this out.
*/