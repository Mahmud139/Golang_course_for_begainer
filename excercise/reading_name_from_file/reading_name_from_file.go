//this programme is from Map chapter of Head first Go book from page no: 337.
// counting votes using Slice
package main

import (
	"fmt"
	"log"
	"golang_tutorial/excercise/datafile"
)

func main(){
	lines, err := datafile.GetStrings("vote.txt")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(lines)
	var names []string
	var counts []int
	for _, line := range lines {
		matched := 1
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = 2
			}
		} 
		if matched == 1 {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}

	for i, name := range names {
		fmt.Printf("%s: %d\n",name, counts[i])
	}
}