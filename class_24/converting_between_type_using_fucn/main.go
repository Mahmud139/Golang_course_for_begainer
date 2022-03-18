//converting_between_type_using_fucn

package main

import "fmt"

type Litters float64
type Gallons float64

func ToLitters(g Gallons) Litters {
	return Litters(g * 3.785)
}

func ToGallons(l Litters) Gallons {
	return Gallons(l * 0.264)
}

func main(){
	carFuel := Gallons(1.2)
	busFuel := Litters(4.5)
	carFuel += ToGallons(Litters(40.0))
	busFuel += ToLitters(Gallons(30.0))
	fmt.Println(carFuel)
	fmt.Println(busFuel)
}