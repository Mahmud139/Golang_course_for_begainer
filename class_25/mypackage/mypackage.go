package mypackage

import "fmt"

type EmbeddedType string

type MyType struct {
	EmbeddedType
}

func (e EmbeddedType) ExportedMethod() {
	fmt.Println("Hi from ExportedMethod on EmbeddedType")
}

func (e EmbeddedType) unexportedmethod() {
	fmt.Println("Hii from unexportedmethod on EmbeddedType")
}