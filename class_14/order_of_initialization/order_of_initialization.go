/* in the same package, if the initialization of anything in "a" depends on 
things in "b" they will be initialized first.*/
package main

import "fmt"

var name = myName()

/* this function will execute first, since name variable's value is depend on this function.*/
func myName() string { 
	fmt.Println("this is the first statement.")
	return "Mahmudul Hasan"
}

/* this init function should be initialized first, but inside the init function the name 
variable's value depands on myName() function. so the myName function will be initialized 
first. */
func init(){ 
	fmt.Println(name)
	name = "Shihab Uddin"
	fmt.Println(name)
}

func main(){
	fmt.Println("my name is",name)
}

