//The panic value is returned from recover
package main

import "fmt"

func main(){
	defer calmDown()
	panic("oh no!")
}

func calmDown(){
	fmt.Println(recover())
	//recover()
}
/*when there is a panic, recover returns whatever value was passed to panic.
This can be used to gather information about the panic, to aid in recovering or to
report errors to the user.*/