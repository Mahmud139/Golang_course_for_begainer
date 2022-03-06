package main

import "fmt"

func main(){
	// var nilMap map[int]string
	// fmt.Printf("%#v\n",nilMap)
	// nilMap[3] = "three"

	var myMap map[int]string = make(map[int]string)
	myMap[3] = "three"
	fmt.Printf("%#v\n",myMap)
}