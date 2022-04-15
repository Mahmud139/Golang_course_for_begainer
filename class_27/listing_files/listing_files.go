package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	files, err := ioutil.ReadDir("M:/code_of_Golang/go_workspace/src/golang_tutorial/class_27/my_directory")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Directory:",file.Name())
		} else {
			fmt.Println("File:",file.Name())
		}
	}
}