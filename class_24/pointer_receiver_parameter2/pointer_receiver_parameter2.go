package main

import "fmt"

type MyType string

func (m MyType) valueMethod() {
	fmt.Println("Method with value receiver")
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}

func main() {
	value := MyType("a value") //a variable with a nonpointer type
	pointer := &value // a variable with pointer type
	value.valueMethod()
	value.pointerMethod() //Go will automatically convert the receiver to a pointer for you

	pointer.valueMethod() //Go will automatically get the value at the pointer for you
	pointer.pointerMethod()
}

/*Notice that we didn’t have to change the method call at all. When you call a
method that requires a pointer receiver on a variable with a nonpointer type, Go
will automatically convert the receiver to a pointer for you. The same is true for
variables with pointer types; if you call a method requiring a value receiver, Go
will automatically get the value at the pointer for you and pass that to the
method.*/