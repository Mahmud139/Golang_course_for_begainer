package main

import "fmt"

func main(){
	//find the avertage value using array
	bags := [4]float64{12,23,34,11}
	var sum float64
	for i:=0; i < len(bags); i++ {
		sum += bags[i]
	}
	fmt.Println("avarage value:",sum/float64(len(bags))) 
}