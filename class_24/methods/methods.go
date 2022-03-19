package main

import "fmt"

type Litters float64
type Mililitter float64
type Gallons float64

func main(){

}

func ToGallons(l Litters) Gallons {
	return Gallons(l * 0.264)
}

func ToGallons(m Mililitter) Gallons {
	return Gallons(m * 0.000264)
}

//we can’t have two ToGallons functions in the same package!