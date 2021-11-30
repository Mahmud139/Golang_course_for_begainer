package main

import "fmt"

func main() {
	notes := [3]string{"do", "or", "die"}
	for index, notes := range notes {
		fmt.Println(index, notes)
	}

	primes := [3]int{3, 5, 7}
	for index, prime := range primes{
		fmt.Println(index, prime)
	}

	// we can omit both index, and the value 
	// i := 0
	// prime := [3]int{3, 5, 7}
	// for range prime {
	// 	fmt.Println(prime[i])
	// 	i++
	// }
}