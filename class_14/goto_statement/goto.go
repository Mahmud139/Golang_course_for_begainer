package main 

import "fmt"

func main(){
	// fmt.Println("a")
	// goto FINISH
	// //fmt.Println("b")
	// FINISH: 
	// 	fmt.Println("c")

	learnGoto()
}

func learnGoto(){
	fmt.Println("a")
	goto FINISH
	//fmt.Println("b")
	FINISH: 
		fmt.Println("c")
}

// func test() {
// 	FINISH:
// 		fmt.Println("c")
// }