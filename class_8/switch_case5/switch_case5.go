//Both switch statement and switch expression are absent.
package main

import "fmt"

func main(){
	age := 45
    switch {
    case age < 18:
        fmt.Println("Kid")
    case age >= 18 && age < 40:
        fmt.Println("Young")
    default:
        fmt.Println("Old")
    }
}

/*If the switch expression is not provided then the default type assumed by
the compiler is boolean. Then the compiler assume that the expression is true.*/

/*The type of switch expression and all of case expression* should match otherwise 
there will be a compiler error raised*/