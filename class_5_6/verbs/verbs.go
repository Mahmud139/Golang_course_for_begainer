package main

import "fmt"

// type point struct {
//     x, y int
// }

func main(){
	//p := point{1, 2}
	// problem we faced before verbs come...
	chinese:= 89
	math:= 90
	fmt.Println("Shihab, You get  on chinese and  on math.",chinese, math)
	/* if we run this program, our desire output won't come out...that's our problem*/

	//to solve this problem we need to use "verbs"..
	//fmt.Printf("Shihab, You get %d on chinese and %d math.",chinese, math)

	// %v verb
	//fmt.Printf("my age is %v",chinese)

	// %+v verb
	//fmt.Printf("struct2: %+v\n", p)
	/*If the value is a struct, the %+v variant will include the struct’s field names.*/

	// %#v verb
	// for more understanding pleas look at the picture at class_5 folder
	//fmt.Printf("my age is %#v\n",chinese)
	/* The %#v variant prints a Go syntax representation of the value, 
	i.e. the source code snippet that would produce that value. */

	// %T verb
	//fmt.Printf("type: %T\n", chinese) //To print the type of a value, use %T.
	

	// %% verb
	//fmt.Printf("i got 90%% marks on chinese listening test.\n")
	/* a literal percent sign; consumes no value */

	// %t verb
	//fmt.Printf("bool: %t\n", true) //Formatting booleans is straight-forward.

	// %d verb
	//fmt.Printf("int: %d\n", 123) 
	/* There are many options for formatting integers. Use %d for standard, base-10 formatting.*/

	// %f verb
	//fmt.Printf("float1: %f\n", 78.9)
	/*There are also several formatting options for floats. For basic decimal formatting use %f.*/

	// %s verb
	//fmt.Printf("string: %s\n", "mahmud") //For basic string printing use %s.
	
	// %b verb
	//fmt.Printf("binary: %b\n", 255) //This prints a binary representation.

	// %o verb
	//fmt.Printf("octal: %o\n", 123)

	// %c
	//fmt.Printf("char: %c\n", 6512)
	/* This prints the character corresponding to the given integer. 
	the character represented by the corresponding Unicode code point*/

	// %q
	//fmt.Printf("double_quotes: %q\n", 6512) //To double-quote strings as in Go source, use %q.

	// %x
	//fmt.Printf("hex: %x\n", 456) //%x provides hex encoding.

	//var name string = "mahmud"
	// %p
	///fmt.Printf("pointer: %p\n", &p) //To print a representation of a pointer, use %p.


	// more formating on integer number
	// fmt.Printf("width_int: |%d|%d|\n", 12, 345)
	// fmt.Printf("width_int: |%6d|%6d|\n", 12, 345)

	// restrict the decimal precision of floating point number
	//fmt.Printf("precision: %.2f\n", 78.912121)

	//To left-justify, use the - flag.
	// fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.212121, 3.412125)
	// fmt.Printf("width3: %-6.2f\n", 1.2)
	// fmt.Printf("width3: %.2f\n", 1.21212)

	
	/*You may also want to control width when formatting strings, especially to ensure that 
	they align in table-like output. For basic right-justified width.*/
	// fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
	// fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")


}
