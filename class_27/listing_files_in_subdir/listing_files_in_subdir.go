package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readdir(file os.FileInfo) {
	f, e := ioutil.ReadDir(file.Name())
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(f)
	// for _, fi := range f {
	// 	fmt.Println(fi.Name())
	// }
	
}

func main() {
	files, err := ioutil.ReadDir("M:/code_of_Golang/go_workspace/src/golang_tutorial/class_27/my_directory")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			readdir(file)
			//fmt.Println("Directory:",f.Name())
			//fmt.Println("Directory:",file.Name())
		} else {
			fmt.Println("File:",file.Name())
		}
	}
}