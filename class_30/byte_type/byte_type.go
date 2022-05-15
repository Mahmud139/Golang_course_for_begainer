package main

import "fmt"

func main() {
	bs := string([]byte{72, 101, 108, 108, 111})
	fmt.Println(bs)

	sb := []byte("Hello")
	fmt.Println(sb)
}