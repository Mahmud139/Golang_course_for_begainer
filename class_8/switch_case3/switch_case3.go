/*Only switch statement is present.Semicolon needs to be 
present after switch statement.*/
package main

import "fmt"

func main(){
	//age := 23
	switch age := 23; {  //Notice the ';' after the statement
	case age < 18 :
		fmt.Println("Kid")
	case age >= 18 && age <= 30 :
		fmt.Println("Young")
	case age > 30 && age <= 50 :
		fmt.Println("Middle age")
	// case 23 :
	// 	fmt.Println("mahmud")
	default:
		fmt.Println("Old")
	}
}

/*When switch expression is not provided the type of all 
case expression* needs to be a boolean too.*/

/*The type of switch expression and all of case expression* should match otherwise 
there will be a compiler error raised*/