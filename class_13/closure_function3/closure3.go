/*Every closure is bound to its own surrounding variable. 
Let's understand what this means by using a simple example.*/

package main

import "fmt"

func main(){
	a := appendStr() //the function appendStr returns a closure.
	fmt.Println(a("World"))
	fmt.Println(a("from Mahmud"))

	b := appendStr()
	fmt.Println(b("Shihab"))
	fmt.Println(b("Uddin"))
}

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string { //This closure is bound to the variable t.
		t = t + " " + b
		return t
	}
	return c
}
