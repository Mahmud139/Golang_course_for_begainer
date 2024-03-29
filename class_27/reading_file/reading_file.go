package main

import (
	"fmt"
	"bufio"
	"log"
	"strconv"
	"os"
)

func OpenFile(fileName string) (*os.File, error){
	fmt.Println("Opening", fileName)
	return os.Open(fileName)
}

func CloseFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}

func GetFloats(fileName string) ([]float64, error){
	var numbers []float64
	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	defer CloseFile(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(),64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	//CloseFile(file)
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}

func main(){
	// numbers, err := GetFloats(os.Args[1])
	numbers, err := GetFloats("sum.txt") //run this file in command prompt
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Sum:%0.2f\n",sum)
}