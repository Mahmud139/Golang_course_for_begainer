package main

import "fmt"

func main(){
	count(1,3)
}

func count(start, end int) {
	fmt.Printf("count(%d, %d) called\n", start, end)
	fmt.Println(start)
	if start < end {
		count(start+1, end)
	}
	fmt.Printf("Returning from count(%d, %d) call\n",start, end)
}

// please use the "recursion2.png" picture to demonstrate how this function work

// func count(start, end int) {
// 	fmt.Println(start)
// 	if start < end {
// 		count(start+1, end)
// 	}
// }