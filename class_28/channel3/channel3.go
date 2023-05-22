//declaraing channel
package main

import (
	"fmt"
	"reflect"
)

func main(){
	// var myChannel chan float64
	// myChannel = make(chan float64)

	myChannel := make(chan float64)
	// myChannel <- 12.1
	// fmt.Println(<- myChannel)
	fmt.Println(reflect.TypeOf(myChannel).Kind())
}