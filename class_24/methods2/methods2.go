package main

import "fmt"

type MyType string

func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}

func main(){
	value := MyType("a MyType value")
	value.sayHi()
	anotherValue := MyType("another value")
	anotherValue.sayHi()
}

/* Once a method is defined on a type, it can be called on any value of that type.
Here, we create two different MyType values, and call sayHi on each of them.*/

/*The name of the receiver parameter in the method definition isn’t important, but its type is; the method you’re defining becomes associated with "all values" of that type.*/