//this programme is from head first go book page 90
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*type a number and press Enter, and we’ll store the number they typed in a variable.*/
	fmt.Print("Enter a number: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
	fmt.Printf("%T", input)
}

/*Next, we need a way to read (receive and store) input from the program’s
standard input, which all keyboard input goes to. The line reader :=
bufio.NewReader(os.Stdin) stores a bufio.Reader in the reader variable
that can do that for us.

To actually get the user’s input, we call the ReadString method on the Reader.
The ReadString method requires an argument with a rune (character) that marks
the end of the input. We want to read everything the user types up until they
press Enter, so we give ReadString a newline rune.*/
