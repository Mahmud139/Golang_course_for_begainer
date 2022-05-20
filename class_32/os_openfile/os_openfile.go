package main

import (
	// "bufio"
	// "fmt"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// file, err := os.OpenFile("tests.txt", os.O_RDONLY, os.FileMode(0600))
	// check(err)
	// defer file.Close()
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
	// check(scanner.Err())

	// file, err := os.OpenFile("tests.txt", os.O_WRONLY, os.FileMode(0600))
	// check(err)
	// _, err = file.Write([]byte("Mahmud\n"))
	// check(err)
	// err = file.Close()
	// check(err)

	// options := os.O_WRONLY | os.O_APPEND
	// file, err := os.OpenFile("tests.txt", options, os.FileMode(0600))
	// check(err)
	// _, err = file.Write([]byte("IS My name\n"))
	// check(err)
	// err = file.Close()
	// check(err)

	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("tests2.txt", options, os.FileMode(0600))
	check(err)
	_, err = file.Write([]byte("Mahmud is My name\n"))
	check(err)
	err = file.Close()
	check(err)
}