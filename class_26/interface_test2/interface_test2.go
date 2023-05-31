package main

import "fmt"

type ReadCloser interface {
	Reader
	Closer
}

type Reader interface {
	Read(string)
}

type Closer interface {
	Close()
}

type BodyEOFSignal struct {
	body ReadCloser
}

func (b *BodyEOFSignal) Read(m string) {
	b.body.Read("world")
	fmt.Println(m)
}

func (b *BodyEOFSignal) Close() {
	b.body.Close()
	fmt.Println("nothing")
}

type Response struct {
	Body ReadCloser
	Age  int
}

func main() {
	rs := Get("hello")

	rs.Body.Read("world")
	rs.Body.Close()
}

// func Get(url string) *Response {
// 	return &Response{}
// }

func Get(url string) Response {
	return Response{}
}