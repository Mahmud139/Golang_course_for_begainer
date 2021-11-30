// multiple init function in same file 
package main

import "fmt"

var initCounter int

func init(){
	fmt.Println("Printed from first init() function.")
	initCounter++
}

func init(){
	fmt.Println("Printed from second init() function.")
	initCounter++
}

func init(){
	fmt.Println("Printed from 3rd init() function.")
	initCounter++
}

func main(){
	fmt.Println("init counter:", initCounter)
}