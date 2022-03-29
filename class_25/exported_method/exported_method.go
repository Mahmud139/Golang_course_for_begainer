package main

import (
	"golang_tutorial/class_25/mypackage"
)

func main(){
	value := mypackage.MyType{}
	/*Because EmbeddedType defines an exported method (named ExportedMethod),
	that method is promoted to MyType, and can be called on MyType values.*/
	value.ExportedMethod()

	/*As with unexported fields, unexported methods are not promoted. You’ll get an
	error if you try to call one.*/
	//value.unexportedMethod()
}