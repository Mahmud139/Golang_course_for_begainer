package main

import (
	"golang_tutorial/class_29/prose"
	"fmt"
)

func main(){
	fruits := []string{"mango", "apple", "pear", "banana"}
	fmt.Println(prose.JoinWithCommas(fruits))
	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("a photo of", prose.JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "prize bull"}
	fmt.Println("a photo of", prose.JoinWithCommas(phrases))
}