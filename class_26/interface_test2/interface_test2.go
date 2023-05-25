package main

import "fmt"

type ResponseWriter interface {
	Write([]byte) int
}

type response struct {
	status bool
	ResponseWriter
}

func (r response) Write(b []byte) int {
	return 12
}

func main() {
	var rs ResponseWriter
	value := rs.Write([]byte("hello"))
	fmt.Println(value)
}