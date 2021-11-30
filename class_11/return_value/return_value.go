package main

import "fmt"

func main(){
	fmt.Println(status(60.1))
	fmt.Println(status(59))
}

func status(grade float64) string {
	if grade < 60.00 {
		return "failing"
	}

	return "passing"
}