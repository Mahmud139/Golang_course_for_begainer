package main

import "fmt"

func main(){
	welcome := []string{"Hello", "World"}
	fmt.Println("before passing to change(): ", welcome)
	change(welcome...)
	fmt.Println("after passing to change(): ", welcome)
}

func change(s ...string) {
	//scenario 1
	s[0] = "Go"
	fmt.Println("inside change(): ", s)

	//scenario 2
	// s = append(s, "Nice to meet you")
	// fmt.Println("inside change(): ", s)

	//scenario 3
	// s = append(s, "Nice to meet you")
	// s[0] = "Go"
	// fmt.Println("inside change(): ", s)
}

/*if you can't understand what's happening here then please study "slice" class again*/