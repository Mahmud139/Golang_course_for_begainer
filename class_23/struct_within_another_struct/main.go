//struct_within_another_struct

package main

import (
	"fmt"
	"golang_tutorial/excercise/magazine"
)

func main(){
	address := magazine.Address{
		Street: "gabtoli",
		City: "Narsingdi",
		State: "NF",
		PostalCode: "1212",
	}

	subscriber := magazine.Subscriber{Name: "mahmud"}
	subscriber.HomeAddress = address
	fmt.Println(subscriber.HomeAddress)
}