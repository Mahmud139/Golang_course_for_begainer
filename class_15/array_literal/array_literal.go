package main

import "fmt"

func main(){
	// Array literal
	var notes [3]string = [3]string{0: "do",2: "or"}
	//var notes [3]string = [3]string{"do","or"}
	fmt.Println(notes[0])
	fmt.Println(notes[1])
	fmt.Println(notes[2])


	// var primes [3]int = [3]int{3,7}
	// fmt.Println(primes[0])
	// fmt.Println(primes[1])
	// fmt.Println(primes[2])

	// we can also use array literal using short hand declaration with := 
	// notes := [3]string{"do", "or", "die"}
	// primes := [3]int{3,5,7}
	// fmt.Println(notes[0])
	// fmt.Println(primes[0])

	text := [3]string{
		"My name is Mahmudul Hasan",
		"My friends name is Shihab Uddin",
		"together we can go far, InshaAllah",
	}
	fmt.Println(text[2])
}

