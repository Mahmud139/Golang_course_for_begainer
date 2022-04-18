/*If you call recover when a program is panicking, it will stop the panic. But
when you call panic in a function, that function stops executing. So there’s no
point calling recover in the same function as panic, because the panic will
continue anyway:*/

package main

import "fmt"

func main(){
	freckOut()
	fmt.Println("Exiting Normally")
}

func freckOut() {
	panic("Oh no!")
	//recover()
}