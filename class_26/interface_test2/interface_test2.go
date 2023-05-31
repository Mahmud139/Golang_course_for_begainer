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
	reader ReadCloser
}

func (b *BodyEOFSignal) Read(m string) {
	b.reader.Read("world")
	fmt.Println(m)
}

func (b *BodyEOFSignal) Close() {
	b.reader.Close()
	fmt.Println("nothing")
}

type Response struct {
	Body ReadCloser
	Age  int
}

func main() {
	rs := Get("hello")

	body := &BodyEOFSignal{
		reader: rs.Body,
	}

	body.Read("world")
	body.Close()
}

func Get(url string) *Response {
	return &Response{Age: 12, Body: nil}
}

// func Get(url string) Response {
// 	return Response{}
// }