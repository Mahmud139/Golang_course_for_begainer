package main

import "fmt"

// func init(){
// 	fmt.Println("Printed from init() function.")
// }

//we can also initialize a variable 
var name = "Mahmudul Hasan"
func init(){
	fmt.Println("Printed from init() function.")
	name = "Shihab"
	fmt.Println(name)
}

func main(){
	fmt.Println("Printed from main() function.")

	fmt.Printf("my name is %s ",name)
}

/*Notice in this above example, we’ve not explicitly called the init() function 
	anywhere within our program. Go handles the execution for us implicitly. */