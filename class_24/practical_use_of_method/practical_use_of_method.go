package main

import "fmt"

type Litters float64
type Mililiter float64
type Gallons float64

func main() {
	soda := Litters(2)
	fmt.Printf("%0.2f litters equals %0.3f gallons.\n",soda, soda.ToGallons())
	water := Mililiter(500)
	fmt.Printf("%0.2f Mililiter equals %0.3f gallons.\n",water, water.ToGallons())
}

//unlike functions, method names don’t have to be unique, as long as they’re defined on different types.
func (l Litters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Mililiter) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}