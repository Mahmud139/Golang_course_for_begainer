//this programme is from Map chapter of Head first Go book from page no: 354.
// counting votes using Map
package main

import (
	"fmt"
	"log"
	"golang_tutorial/excercise/datafile"
)

// func main(){
// 	lines, err := datafile.GetStrings("vote.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	counts := make(map[string]int)
// 	for _, line := range lines {
// 		counts[line]++
// 	}
// 	fmt.Println(counts)
// }

//updated version of counting vote programme 
func main(){
	lines, err := datafile.GetStrings("vote.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	for key, value := range counts {
		fmt.Printf("Votes for %s : %d\n",key,value)
	}
}