package main

import "fmt"

func main(){
	var mystructs struct {
		number float64
		word	string 
		toggle bool 
	}

	fmt.Printf("%#v\n",mystructs)
}