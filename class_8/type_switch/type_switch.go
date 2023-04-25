package main

import "fmt"

func main(){
	var shihab interface{} = 12
	
	switch a:= shihab.(type) {
	case bool :
		fmt.Println("value is boolean type.")
	case int :
		fmt.Println("value is integer type.")
		fmt.Printf("%T", a)
	case float64 :
		fmt.Println("value is float64 type.")
	default :
		fmt.Printf("value is of type: %T",a)
	}

	//we cann't use this kinda statemaent outside of "type switch"
	// n := shihab.(type)
	// fmt.Printf("%T", n)
}

