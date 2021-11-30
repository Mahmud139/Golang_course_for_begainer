package main

import "fmt"

func main(){
	letters := []string{"a", "b", "c"}

	for i:= 0; i< len(letters); i++ {
		fmt.Println(letters[i])
	}

	for _, value := range letters {
		fmt.Println(value)
	}
}