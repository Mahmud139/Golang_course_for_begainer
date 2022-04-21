package main

//import "fmt"

func main() {
	//fmt.Println(go greetings())
	go greetings()
}

func greetings() string {
	return "Hi"
}

/*Go won’t let you use the return value from a function called with a go statement, 
because there’s no guarantee the return value will be ready before we attempt to use it:
*/