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

	/*You can only get pointers to values that are stored in variables. If you try to
	get the address of a value that’s not stored in a variable, you’ll get an error:*/
	// &MyType("a value").valueMethod()
	// &MyType("a value").pointerMethod()

	/*The same limitation applies when calling methods with pointer receivers. Go
	can automatically convert values to pointers for you, but only if the receiver
	value is stored in a variable. If you try to call a method on the value itself,
	Go won’t be able to get a pointer, and you’ll get a similar error:*/
	// MyType("a value").pointerMethod()
}

/*Notice that we didn’t have to change the method call at all. When you call a
method that requires a pointer receiver on a variable with a nonpointer type, Go
will automatically convert the receiver to a pointer for you. The same is true for
variables with pointer types; if you call a method requiring a value receiver, Go
will automatically get the value at the pointer for you and pass that to the
method.*/