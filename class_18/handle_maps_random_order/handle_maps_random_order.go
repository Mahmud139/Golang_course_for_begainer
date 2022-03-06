package main

import (
	"fmt"
	"sort"
)

func main(){
	grades := map[string]int{"Alma": 12, "Rohit": 34, "Carl": 54}
	var names []string
	for name := range grades {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s has grade of %d%%\n",name, grades[name])
	}
}