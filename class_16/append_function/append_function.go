package main 

import "fmt"

func main(){
	//slide no 22
	slice := []string{"a", "b"}
	fmt.Println(slice, len(slice))
	slice = append(slice, "c")
	fmt.Println(slice, len(slice))
	slice = append(slice, "d", "c")
	fmt.Println(slice, len(slice))
}