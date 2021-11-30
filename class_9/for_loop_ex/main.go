/*Let's modify the program we just wrote so that instead of just printing 
the numbers 1-10 on each line it also specifies whether or not the number is even or odd.*/

package main

import "fmt"

func main(){
	for i:= 1; i<= 10; i++{
		if i%2 == 0 {
			fmt.Printf("%d  even\n",i)
			//fmt.Println(i, "even")
		} else{
			fmt.Printf("%d  odd\n",i)
			//fmt.Println(i, "odd")
		}
	}
}

/* Home Work: Write a program that prints out all the numbers evenly divisible by 3 
between 1 and 100. (3, 6, 9, etc.)*/

/*Write a program that prints the numbers from 1
to 100. But for multiples of three print "Fizz" instead
of the number and for the multiples of five
print "Buzz". For numbers which are multiples
of both three and five print "FizzBuzz".*/