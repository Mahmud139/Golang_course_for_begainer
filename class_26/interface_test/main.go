package main

import "fmt"

type RedCloser interface {
	Reder
	Clser
}

type Reder interface {
	Read(string)
}

type Clser interface {
	Clse()
}

type Respnse struct {
	Body RedCloser
	Age  int
}

func (r *Respnse) Read(b string) {
	fmt.Println(b)
}

func (r *Respnse) Clse() {
	fmt.Println("do nothing")
}

func main() {
	rs := &Respnse{}
	rs.Read("world")
	rs.Body.Read("world")
	rs.Body.Clse()
}

// func Get(url string) *Response {
// 	return &Response{}
// }

// func Get(url string) Response {
// 	return Response{}
// }