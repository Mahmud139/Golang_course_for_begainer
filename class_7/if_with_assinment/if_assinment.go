package main

import "fmt"

func main(){
	if num:= 10; num%2 == 0 {  //num is initialized in the if statement
		fmt.Println(num,"is even number.")
	} else {
		fmt.Println(num, "is odd number.")
	}
}



/*One thing to be noted is that num is available only for access from inside 
the if and else. i.e. the scope of num is limited to the if else blocks. 
If we try to access num from outside the if or else, the compiler will complain.*/