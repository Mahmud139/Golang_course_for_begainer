package main

import "fmt"

type testError string
func (t testError) Error() string {
	return string(t)
}

type shihabError string
func (s shihabError) Error() string {
	return string(s)
}

// type error interface {
//     Error() string
// }

func main() {
	var err error = testError("what's your name?")
	fmt.Println(err.Error())

	var err2 error = shihabError("your name?")
	fmt.Println(err2.Error())
}