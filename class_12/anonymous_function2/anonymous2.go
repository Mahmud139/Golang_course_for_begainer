// anonymous function assigning to a variable 
package main

import "fmt"

func main(){
	//with parameter and no return value 
	// area := func(height int, width int)  {
	// 	fmt.Println("The area is", height*width)
	// }
	// area(12,12)
	// fmt.Printf("The type of a() is: %T", area)

	//with parameter and return value
	area := func(height int, width int) int {
		return height*width
	}
	//area(12,12) //in this time we can't see the output value in our screen 
				// so we need to assign this variable in our Println function. 
	//fmt.Println(area(12,12))
	res := area(12,12)
	fmt.Println(res)
	fmt.Printf("The type of a() is: %T", area)

	// no paremeter but have return value (home work)


}