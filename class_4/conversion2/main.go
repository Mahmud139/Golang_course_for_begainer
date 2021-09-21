package main

import "fmt"

func main(){
	var length float64 = 1.2
	var width int = 2
	
	fmt.Println("Area is", length*float64(width)) /*convert the int to float64 before 
	multiplying with another float64. */
	fmt.Println("length > width?", length > float64(width)) /* convert the int to float64 
	before comparing with another float64 */

	/* Your task: Now let’s try converting an int to a float64 before assigning 
	it to a float64 variable */

}