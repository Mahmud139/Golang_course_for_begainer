package main

import "fmt"

func main(){
	res:= foo()
	fmt.Println(res)
}

func foo() (result string) {
	defer func ()  {
		result = "change world"
	}()
	return "hello world"
}