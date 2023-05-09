package main

import (
	"fmt"
)

func main() {
	var ranks map[string]int
	ranks = make(map[string]int)

	//name := make(map[string]int) //short-hand declaration

	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Println(ranks["gold"])
	fmt.Println(ranks["silver"])
	fmt.Println(ranks)


	//in map the key is Alpha-numerically ordered
	// order := make(map[string]int)
	// order["b"] = 2
	// order["a"] = 1
	// order["c"] = 3
	// fmt.Println(order)
	// fmt.Printf("%#v\n", order)

	// order2 := make(map[int]int)
	// order2[2] = 2
	// order2[1] = 11
	// order2[3] = 33
	// fmt.Println(order2)
}