package main

import "fmt"

func main(){
	//slide no 34
	//lets declare a slice with sort-declaration operator
	// n := []int{1,2,3}
	// fmt.Println(len(n)) //3
	// fmt.Println(cap(n)) //3

	// n = append(n, 12,23)
	// fmt.Println(len(n)) //5
	// fmt.Println(cap(n)) //6

	/* in the above example, we have seen that when we add elements to our slice beyond it's
	capacity, the capacity increased by double. this is because of when we add elements to 
	our slice beyond it's capacity using append function, all its elements of our previous 
	slice will be copied to a new, larger array, and the slice will be updated to refer to
	this new array. and most importantly the larger array's size is doubled than previous array.
	*/

	// now guess what will be the capacity of slice after appending two elements
	// n = append(n, 45,26)
	// fmt.Println(len(n))
	// fmt.Println(cap(n))

	/* just imagine if we make a programme that dynamically append elements to our slice, that
	will occur unesseccary memory allocation for our programme. and this kinda bug will makes
	our programme slower.  */

	//********************** IMPORTANT NOTE ************************
	/* When slice length is less than capacity, using the append function, the length of the 
	slice will be increased by one without any change in its capacity. Let's see a example */

	x := make([]int, 3, 5)
	x[0] = 1
    x[1] = 2
    x[2] = 3
	fmt.Printf("x=%v\n", x)
    fmt.Printf("length=%d\n", len(x))
    fmt.Printf("capacity=%d\n", cap(x))

	//Append number 4
    x = append(x, 4)
    fmt.Println("\nAppend Number 4")
    fmt.Printf("x=%v\n", x)
    fmt.Printf("length=%d\n", len(x))
    fmt.Printf("capacity=%d\n", cap(x))

    //Append number 5
    x = append(x, 4)
	fmt.Println("\nAppend Number 4")
    fmt.Printf("x=%v\n", x)
    fmt.Printf("length=%d\n", len(x))
    fmt.Printf("capacity=%d\n", cap(x))
	
	/*Capacity in all cases doesn't changes and it is 5 while length increases by 1.*/

	/*Scenario 2 : When slice length is equal than capacity : In this case since there 
	is no more capacity, so no new elements can be accommodated.  So in this case under 
	the hood an array of double the capacity will be allocated. The current array pointed 
	by the  slice will be copied to that new array. Now the slice will starting pointing 
	to this new array. Hence the capacity will be doubled and length will be increased by 1.*/






	/*So the Question is how we can prevent this kinda problem while we need to dynamically 
	appending elements to our slice ? */

	/*Let’s take the scenario where we want to make a list of numbers, 0 through 3, and all the 
	elements are dynamically allocated. we can solve this in two ways:*/

	//first method: use append() in a loop
	numbers := []int{} 
	fmt.Println(len(numbers), cap(numbers))
	for i := 0; i < 4; i++ {
		numbers = append(numbers, i)
		//numbers[i] = i  // can we use this statement instead of using append function?
		// we have discussed about it on slide no 25
	} 
	fmt.Println(numbers)
	fmt.Println(len(numbers), cap(numbers))
	/* the above programme makes our programme slower. because of When adding to an empty slice, 
	each time you make a call to append, the program checks the capacity of the slice. If the 
	added element makes the slice exceed this capacity, the program will allocate additional 
	memory(doubled from previous array) to account for it. This creates additional overhead 
	in your program and can result in a slower execution. */


	/* second method: we can pre-allocate the slice first and use cap() to loop through to 
	fill the values.*/ 
	// num := make([]int, 4)
	//fmt.Println(len(num), cap(num))
	// for i := 0; i < cap(num); i++ {
	// 	num[i] = i
	// 	//num = append(num, i) // can we use this statement instead of using above statement?
	// }
	// fmt.Println(num)
	//fmt.Println(len(num), cap(num))


	// x := []int{} //we cant use this statement with line no 71 statement
	// //x := make([]int, 4) // we cant use this statement with append function 
	// for i := 0; i < 4; i++ {
	// 	x = append(x, i)
	// 	//numbers[i] = i 
	// } 
	// fmt.Println(x)


	/*While the append() and the cap() strategies are both functionally equivalent, 
	the cap() example avoids any additional memory allocations that would have been 
	needed using the append() function.*/

}