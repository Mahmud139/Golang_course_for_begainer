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

type Response struct {
	Body ReadCloser
	Age  int
}

func (r *Response) Read(b string) {
	fmt.Println(r.Body)
}

func (r *Response) Close() {
	fmt.Println("do nothing")
}

func main() {
	res := Get("hello")
	res.Body.Read("world")
	res.Body.Close()

}

func Get(url string) *Response {
	return &Response{}
}