package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func scanDirectory(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("returning error from scandirectory111 (\"%s\") call\n", path)
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err = scanDirectory(filePath)
			if err != nil {
				fmt.Printf("returning error from scandirectory (\"%s\") call\n", path)
				return err
			}
		} else {
			fmt.Println(filePath)
		}
	}
	return nil
}

func main() {
	err := scanDirectory("M:/code_of_Golang/go_workspace/src/my_directory")
	if err != nil {
		log.Fatal(err)
	}
}

/*  If an error occurs in one of the recursive scanDirectory calls, that error has to
	be returned up the entire chain until it reaches the main function!*/

// recursive function call manei alada alada function call and protita function call er shathe ager function call er
// kono relation nai.
