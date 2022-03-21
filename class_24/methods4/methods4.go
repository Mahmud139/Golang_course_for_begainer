//A method is (pretty much) just like a function
package main

import "fmt"

type MyType string

func (m MyType) MethodWithParameter(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}

func main(){
	value := MyType("MyType Value")
	value.MethodWithParameter(4, true)
}