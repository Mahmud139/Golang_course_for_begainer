// There is an alternative—users could pass the values to the program as
// command-line arguments. Head_first_go 314 page
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main(){
	/* The os package has a package variable, os.Args, that gets set to a slice of
strings representing the command-line arguments the currently running program
was executed with. */
	//fmt.Println(os.Args)

	/* on your terminal type the programme name, but don’t hit the Enter key just yet. 
	Following the program name, type a space, and then type one or more arguments, 
	separated by spaces. Then hit Enter. The program will run and print the value of os.Args.
	"the name of the executable is being included as the first element of os.Args." */

	/*If we use a slice operator of [1:] on os.Args, it will give us a new slice that omits 
	the first element (whose index is 0), and includes the second element (index 1) through
	the end of the slice.*/
	//fmt.Println(os.Args[1:]) //get a new slice that includes the second element through the end of os.Args

	arguments := os.Args[1:]
	var sum float64 = 0
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	samplecount := float64(len(arguments))
	fmt.Printf("Average : %0.2f",sum/samplecount)
}