package main

import "fmt"

func main() {
	var ranks map[string]int
	ranks = make(map[string]int)

	//name := make(map[string]int)

	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Println(ranks["gold"])
	fmt.Println(ranks["silver"])
	fmt.Println(ranks)
}