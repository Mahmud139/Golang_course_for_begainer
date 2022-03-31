package mypkg

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameter called")
}

func (m MyType) MethodWithParameter(f float64) {
	fmt.Println("methodwithparameter called", f)
}

func (m MyType) MethodWithReturnValue() string {
	return "methodWithReturnValue called"
}

func (m MyType) MethodNoInterface() {
	fmt.Println("MethodNoInterface called")
}