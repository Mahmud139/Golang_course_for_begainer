package main

import "fmt"

func main(){
	var shihab interface{} = 12
	
	switch a:= shihab.(type) {
	case bool :
		fmt.Println("value is boolean type.")
	case int :
		fmt.Println("value is integer type.")
	case float64 :
		fmt.Println("value is float64 type.")
	default :
		fmt.Printf("value is of type: %T",a)
	}


}

