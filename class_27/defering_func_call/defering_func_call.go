package main

import (
	"fmt"
	"log"
)

// func Socialize() {
// 	defer fmt.Println("Goodbye")
// 	fmt.Println("Hello!")
// 	fmt.Println("Nice Weather, eh?")
// }

func Socialize() error {
	defer fmt.Println("Goodbye")
	fmt.Println("Hello!")
	return fmt.Errorf("I don't want to talk.")
	// fmt.Println("Nice Weather, eh?")
	// return nil 
}
	/*because we include a defer keyword before the
	fmt.Println("Goodbye!") call, Socialize will always be polite enough to
	print “Goodbye!” before ending the conversation.*/

func main(){
	err := Socialize()
	if err != nil {
		log.Fatal(err)
	}
}