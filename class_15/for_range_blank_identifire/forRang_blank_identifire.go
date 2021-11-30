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

	// var primes [3]int = [3]int{3, 5, 7}
	// for index, _ := range primes{
	// 	fmt.Println(index)
	// }
}