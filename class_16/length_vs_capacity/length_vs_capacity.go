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


	//IMPORTANT NOTE: don't use cap() for conditional part of loop
	n := make([]int, 4)
	fmt.Println(n)
	fmt.Println(cap(n))
	fmt.Println(len(n))
	for i := 0; i < cap(n); i++ { //don't use cap() here with inside append()
		//n[i] = i 
		n = append(n, i)
		/*the above statement is not valid here, because "i < cap(n)" this condition is unreachable with the above statement. every time this statement run, capacity of n is also increasing. every time when this statement run, program try to allocate more memory. at the end we got "running out of memory" error.*/
	} 
	fmt.Println(n)
	fmt.Println(cap(n))
	fmt.Println(len(n))
}