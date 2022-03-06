package main

import "fmt"

func main(){
	//slide no 19
	var ok bool 
	ranks := make(map[string]int)
	var rank int
	ranks["bronze"] = 3
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n",rank, ok)
	delete(ranks, "bronze")
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n",rank, ok)

	isPrime := make(map[int]bool)
	var prime bool 
	isPrime[5] = true
	prime, ok = isPrime[5]
	fmt.Printf("isPrime: %v, ok: %v\n",prime, ok)
	delete(isPrime, 5)
	prime, ok = isPrime[5]
	fmt.Printf("isPrime: %v, ok: %v\n",prime, ok)
}