package main

import "fmt"

func main(){
	//name := "mahmud"
	

	// fmt.Println(len(name)) //6
	// fmt.Println(len(name2)) //17

	// for i:= 0; i<len(name2); i++ {
	// 	fmt.Printf("%c   %v\n", name2[i],name2[i])
	// }
	name2 := "a马姆b"
	// 0 1 4 7 
	// for indx, value := range name2 {
	// 	fmt.Printf("%v      %c\n",indx, value)  // M
	// }

	for indx, value := range name2 {
		fmt.Println(indx, value)
	}

	//a, b, c := 12, 213, 34 

	for _, value := range name2 {
		fmt.Println(value)
	}

	for i := range name2 {
		fmt.Println(i)
	}

	for range name2 {
		fmt.Println("hello")
	}

}