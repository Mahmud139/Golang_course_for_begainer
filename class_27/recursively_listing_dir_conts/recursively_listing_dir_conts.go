package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func scanDirectory(dirName string) {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(files[0])
	fmt.Printf("%T\n",files[0])
	for _, file := range files {
		fmt.Printf("%v\n",file)
		if file.IsDir() {
			fmt.Printf("Opening %s\n",file.Name())
			scanDirectory(dirName + "/" +file.Name()) //need to correction, because the system will not find the file/dir location
		} else {
			fmt.Println("File:",file.Name())
		}
	}
}

func main() {
	scanDirectory("M:/code_of_Golang/go_workspace/src/golang_tutorial/class_27/my_directory")
}