//this programme is from 276 page of Head first go book

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)
// ******* humbly request to me: please read the function description to more understand how the function is work.
func main(){
	file, err := os.Open("data.txt") //open the data file for reading
	if err != nil {
		log.Fatal(err)
	} //if there wan an error opening the file, report it and exit.

	/*We start by passing a string with the name of the file we want to open to the
os.Open function. Two values are returned from os.Open: a pointer to an os.File value 
representing the opened file, and an error value. As we’ve seen
with so many other functions, if the error value is nil it means the file was
opened successfully, but any other value means there was an error. (This could
happen if the file is missing or unreadable.) If that’s the case, we log the error
message and exit the program.*/

	scanner := bufio.NewScanner(file) //create a new scanner for the file
	/*Then we pass the os.File value to the bufio.NewScanner function. That will
return a bufio.Scanner value that reads from the file.*/

	for scanner.Scan() { //read a line from the file
		fmt.Println(scanner.Text()) // print the line
		fmt.Printf("%T\n",scanner.Text()) //scanner.Scan read a line as a "string" 
	} //loops until the end of the file is reached and scanner.Scan return false

	/* "The Scan method on bufio.Scanner is designed to be used as part of a for loop." 
	It will read a single line of text from the file, returning true if it read data
successfully and false if it did not. If Scan is used as the condition on a for
loop, the loop will continue running as long as there is more data to be read.
Once the end of the file is reached (or there’s an error), Scan will return false,
and the loop will exit. */

	/* After calling the Scan method on the bufio.Scanner, "calling the Text method
returns a string with the data that was read." For this program, we simply call
Println within the loop to print each line out. */

	err = file.Close() //close the file to free resources 
	if err != nil {
		log.Fatal(err)
	} // if there was an error closing the file, report it and exit

	/* "Once the loop exits, we’re done with the file. Keeping files open consumes
resources from the operating system, so files should always be closed when a
program is done with them. Calling the Close method on the os.File will accomplish this." 
Like the Open function, the Close method returns an error value, which will be nil unless there 
was a problem. (Unlike Open, Close returns only a single value, as there is no useful value 
for it to return other than the error.)*/

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	} //if there was an error scanning the file, report it and exit

	/* It’s also possible that the bufio.Scanner encountered an error while scanning
through the file. If it did, calling the Err method on the scanner will return that
error, which we log before exiting. */

}