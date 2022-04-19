package main

import "fmt"

func main() {
	defer calmDown()
	err := fmt.Errorf("there's an error")
	panic(err)
}

func calmDown() {
	p := recover()
	fmt.Println(p.Error())
}