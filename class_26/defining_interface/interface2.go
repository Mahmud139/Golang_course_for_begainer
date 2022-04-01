package main 

import (
	"golang_tutorial/class_26/mypkg"
	"fmt"
)

func main(){
	/*A variable declared with an interface type can hold any value whose type
	satisfies that interface. This code declares a value variable with MyInterface as
	its type, then creates a MyType value and assigns it to value. (Which is allowed,
	because MyType satisfies MyInterface.) Then we call all the methods on that
	value that are part of the interface.*/
	
	var value mypkg.MyInterface
	value = mypkg.MyType(6)
	value.MethodWithParameter(12.1)
	value.MethodWithoutParameters()
	fmt.Println(value.MethodWithReturnValue())
}