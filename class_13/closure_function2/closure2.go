// Use of closure functions

/*One way to use closure is by writing a function which returns another function 
which – when called – can generate a sequence of numbers. 
here we will generate all the even numbers using a closure function.*/

package main

import "fmt"

func main(){
	nextEven := evenGenerator() //evenGenerator returns a function which generates even numbers.
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	/*Each time when we call nextEven function, it adds 2 to the local variable 
	which – unlike normal local variables – persist between calls. for example: */

	secondTime := evenGenerator()
	fmt.Println(secondTime())
	fmt.Println(secondTime())
	fmt.Println(secondTime())
	
	// fmt.Println(add())
	// fmt.Println(add())
	// fmt.Println(add())

}

func evenGenerator() func() uint {
	i := uint(0)
	return func() uint {
		i += 2
		return i
	}
}

// var i int = 0
// func add() int {
// 	// i := 0
// 	i = i + 1
// 	return i
// }