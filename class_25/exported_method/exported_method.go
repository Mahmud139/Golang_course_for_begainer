package main

import (
	"golang_tutorial/class_25/mypackage"
)

func main(){
	value := mypackage.MyType{}
	value.ExportedMethod()
}