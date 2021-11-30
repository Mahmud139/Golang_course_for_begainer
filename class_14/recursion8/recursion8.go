// fibonacci series 
package main

import "fmt"

func main(){
	for i:= 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}
	
	return fibonacci(n-1) + fibonacci(n-2)
}

// 0 1 1 2 3 5