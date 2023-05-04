package main

import "fmt"

func main(){
	//slide no 33
	//Let’s create a simple slice with capacity greater than length.
	numbers := make([]int, 3, 5)

	/* Accessing the slice behind its length will result in a run time error 
	“Index out of range”. "It doesn’t matter if the accessed index is within 
	the capacity". So the below line will cause the run time error. we only can
	do this by using of append() function.*/
	// numbers[4] = 5 //this cause index out of range error in runtime
	// fmt.Println(numbers[4])

	/* The length of the slice can be increased up to its capacity by re-slicing. 
	So below re-slice will increase the length from 3 to 5. */
	numbers = numbers[0:5]
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
	numbers[3] = 12
	fmt.Println(numbers)

	/* The length of the slice can also be decreased using re-slicing. So below 
	re-slice will decrease the length from 3 to 2 */
	numbers = numbers[0:2]
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
	// numbers[3] = 12//this will cause index out of range error in the runtime
	// fmt.Println(numbers)

	/* The advantage of having capacity is that array of size capacity can be 
	pre-allocated during the initialization.  This is a performance boost as 
	if more elements are needed to include in this array then space is already 
	allocated for them. */
}