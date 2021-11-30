package main

import "fmt"

func main(){
	//res:= double(12.12)
	//fmt.Println(res)
	
	fmt.Println(double(12.23))
}

func double(number float64) float64 {
	return number * 2
}