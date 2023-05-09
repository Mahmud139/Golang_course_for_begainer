package main

import "fmt"

func main(){
	var employee = map[string]int{
		"Mark": 10, 
		"Sandy": 20,
		"Rocky": 30, 
		"Rajiv": 40, 
		"Kate": 50,
	}
	fmt.Println(employee)
	//Method - 1
	for k := range employee {
		delete(employee, k)
	}
	fmt.Println(employee)
 
	// Method - II
	// employee = make(map[string]int)
	// fmt.Println(employee)
}