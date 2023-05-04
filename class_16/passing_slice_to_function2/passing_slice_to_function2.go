package main

import "fmt"

func change(s []string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println("inside the function: ",s)
	//Q1: why the result is bit different between main() and this function?


	// s = append(s, "playground")
	// s[0] = "Go"
	// fmt.Println("inside the function: ",s)
	//Q2: why the result between main() and this function is completely different?
}

func main() {
	welcome := []string{"hello", "world"}
	fmt.Println("Before calling the function: ", welcome)
	change(welcome)
	fmt.Println("after calling the function: ",welcome)
}

//to understand the above question, you must have to understand how slice work.
