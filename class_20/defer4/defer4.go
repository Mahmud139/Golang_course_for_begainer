package main

import "fmt"

func main(){
	res:= foo()
	fmt.Println(res)

	// res2:= bar()
	// fmt.Println(res2)
}

/*Deferred anonymous functions may access and modify the surrounding function’s "named return" parameters.*/
func foo() (result string) {
	defer func ()  {
		result = "change world"
	}()
	return "hello world"
}


/*Deferred anonymous functions can't access and modify the surrounding function’s return parameters.*/
// func bar() string {
// 	result := "hello world"
// 	defer func ()  {
// 		result = "change world"
// 	}()
// 	return result
// }