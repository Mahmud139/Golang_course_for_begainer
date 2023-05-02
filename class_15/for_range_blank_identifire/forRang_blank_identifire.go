package main

import "fmt"

func main(){
	// var notes [3]string = [3]string{"do", "or", "die"}
	// for index, notes := range notes {
	// 	fmt.Println( notes)
	// }

	// var primes [3]int = [3]int{3, 5, 7}
	// for index, primes := range primes{
	// 	fmt.Println(primes)
	// }
	
	//we use blank_identifire to ignore/discard the index value. 
	var notes [3]string = [3]string{"do", "or", "die"}
	for _, notes := range notes {
		fmt.Println(notes)
	}

	//we can ignore the value part: and we can write it in two way
	var primes [3]int = [3]int{3, 5, 7}
	for index := range primes{ //this way of writing only valid when we ignore the value part
		fmt.Println(index)
	}
	// for index, _ := range primes{
	// 	fmt.Println(index)
	// }
}