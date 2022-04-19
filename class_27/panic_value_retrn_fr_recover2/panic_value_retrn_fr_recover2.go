package main

import "fmt"

func main() {
	defer calmDown()
	err := fmt.Errorf("there's an error")
	panic(err)
}

func calmDown() {
	p := recover()
	//fmt.Println(p.Error())
	/*To call methods or do anything else with the panic value, you’ll need to convert
it back to its underlying type using a type assertion.*/
	err := p.(error)
	fmt.Println(err.Error())
}