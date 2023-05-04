package main 

import "fmt"

func main(){
	//slide no 32
	x := make([]int, 5, 14)
	fmt.Printf("%#v\n",x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// y := make([]int, 5)
	// fmt.Println(len(y))
	// fmt.Println(cap(y))

	// z := []int{1,2,3}
	// fmt.Println(len(z))
	// fmt.Println(cap(z))
}