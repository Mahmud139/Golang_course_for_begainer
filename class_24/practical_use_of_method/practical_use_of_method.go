package main

type Litters float64
type Mililiter float64
type Gallons float64

func main(){

}

//unlike functions, method names don’t have to be unique, as long as they’re defined on different types.
func (l Litters)ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Mililiter) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}