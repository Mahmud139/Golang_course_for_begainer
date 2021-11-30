// functions in the fmt package know how to handle arrays. 
package main

import "fmt"

func main(){
	var notes [3]string = [3]string{"do", "or", "die"}
	var primes [3]int = [3]int{3,5,7}
	fmt.Println(notes)
	fmt.Println(primes)

	// format arrays as they would appear in Go code 
	fmt.Printf("%#v\n",notes)
	fmt.Printf("%#v\n",primes)

	// name := "shihab"
	// fmt.Println("My name is",name)
	// fmt.Println(name, "is my friend")
	
	// fmt.Println("i am", name, "uddin")
	// fmt.Printf("i am %s uddin. i am %s", name, "mahmud")
}