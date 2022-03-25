//adding_validation_to_setter_method
package main

import (
	"fmt"
	"errors"
	"log"
)

type Date struct {
	Year int
	Month int
	Day int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.Year = year 
	return nil
}

func main(){
	date := Date{}
	//err := date.SetYear(0)
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year)
}