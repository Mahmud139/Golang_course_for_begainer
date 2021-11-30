package main

import "fmt"

func main(){
	//slide no 22
	s1 :=[]string{"s1","s1"}
	fmt.Println(s1, len(s1), cap(s1))

	s2 := append(s1, "s2", "s2")
	fmt.Println(s1, len(s2), cap(s2))

	s3 := append(s2, "s3", "s3")
	fmt.Println(s1, len(s3), cap(s3))

	s4 := append(s3, "s4", "s4", "a")
	fmt.Println(s1, len(s4), cap(s4))
	//fmt.Println(s1, s2, s3, s4)
	// s4[0] = "XX"
	// fmt.Println(s1, s2, s3, s4)

	//slide no 23
	// s1 :=[]string{"s1","s1"}
	// fmt.Println(s1, len(s1))
	// s1 = append(s1, "s2", "s2")
	// fmt.Println(s1, len(s1))
	// s1 = append(s1, "s3", "s3")
	// fmt.Println(s1, len(s1))
	// s1 = append(s1, "s4", "s4")
	// fmt.Println(s1, len(s1))
}