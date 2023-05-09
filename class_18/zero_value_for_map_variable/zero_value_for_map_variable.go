package main

import "fmt"

func main(){
	var nilMap map[int]string //nill map
	fmt.Printf("%#v\n",nilMap)
	nilMap[3] = "three" //we can't assign to nill map
	fmt.Println(nilMap)

	//this will work fine
	// var myMap map[int]string = make(map[int]string)
	// myMap[3] = "three"
	// fmt.Printf("%#v\n",myMap)

	//but we can assing to empty map
	// emptyMap := map[string]int{} //empty map
	// fmt.Printf("%#v\n",emptyMap)
	// emptyMap["roll"] = 12
	// fmt.Println(emptyMap)
}