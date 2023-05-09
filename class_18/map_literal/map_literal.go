package main

import "fmt"

func main(){
	mymap := map[string]int{"a":12, "b": 121, "c": 23}
	fmt.Println(mymap)

	ranks := map[string]int{"gold": 1, "silver": 2, "bronze": 3}
	fmt.Println(ranks["gold"])
	fmt.Println(ranks["silver"])

	elements := map[string]string{
		"name" : "mahmudul hasan", 
		"father" : "Sopon",
	}
	fmt.Println(elements)

	emptyMap := map[string]string{}
	fmt.Println(emptyMap)
	fmt.Printf("%#v", emptyMap)
}