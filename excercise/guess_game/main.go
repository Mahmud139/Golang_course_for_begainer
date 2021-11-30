package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	// 1-100
	seconds:= time.Now().Second()
	rand.Seed(int64(seconds))
	target := rand.Intn(101) 
	//fmt.Println(target)

	fmt.Println("I have chosen a random number between 1 to 100.")
	fmt.Println("Can you guess it?")
    var	guesses int
	
	success := false
	for i:= 0;i<10;i++ {
		
		fmt.Println("You have", 10-i , "guesses left.")
		fmt.Print("Make a guess: ")
		fmt.Scan(&guesses)

		if guesses < target {
			fmt.Println("Opps. Your guess was LOW.")
			fmt.Println()
		} else if guesses > target{
			fmt.Println("Opps. Your guess was HIGH.")
			fmt.Println()
		} else {
			success = true
			fmt.Println("Good job! You guessed it.")
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:",target)
	}


}