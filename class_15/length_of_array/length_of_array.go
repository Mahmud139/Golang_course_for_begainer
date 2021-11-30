package main

import "fmt"

func main(){
	// notes := [3]string{"do", "or", "die"}
	// fmt.Println(len(notes))
	// primes := [3]int{3, 5, 7}
	// fmt.Println(len(primes))


	notes := [3]string{"do", "or", "die"}
	for i := 0; i < len(notes); i++ {
		fmt.Println(notes[i])
	}
}