package main

import "fmt"

func main(){
	//slide no 30
	// mySlice := []int{1,2,3,4,5,6}
	// fmt.Println(mySlice, len(mySlice))
	//now we want to remove the 3rd element of mySlice, the index number of 3rd element is 2.
	// mySlice = append(mySlice[:2], mySlice[3:]...)
	// fmt.Println(mySlice, len(mySlice)) //now the 3rd element is no longer in our mySlice 


	//slide no 31
	//delete a range with the same approach
	// mySlice := []int{1,2,3,4,5,6}
	// fmt.Println(mySlice, len(mySlice))

	// mySlice = append(mySlice[:2], mySlice[5:]...)
	// fmt.Println(mySlice, len(mySlice))

	// extra
	// delete random item from a slice 
	mySlice := []int{1,2,3,4,5,6,7,8}
	fmt.Println(mySlice, len(mySlice))

	mySlice = append(mySlice[:2], mySlice[3], mySlice[5],mySlice[6],mySlice[7])
	//mySlice = append(mySlice[:2], mySlice[3:5], mySlice[5:]...)
	//mySlice = append(mySlice, 1,2,3,4,5)
	fmt.Println(mySlice, len(mySlice))

}