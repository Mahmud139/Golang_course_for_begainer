package main

//import "fmt"

type Liters float64
type Mililiter float64
type Gallons float64

func main(){

}

// func ToGallons(l Litters) Gallons {
// 	return Gallons(l * 0.264)
// }

// func ToGallons(m Mililitter) Gallons {
// 	return Gallons(m * 0.000264)
// }
//we can’t have two ToGallons functions in the same package!

func LiterToGallon(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func MililiterToGallon(m Mililiter) Gallons {
	return Gallons(m * 0.000264)
}

func GallonsToLiter(g Gallons) Liters {
	return Liters(g * 3.785)
}

func GallonsToMililiters(g Gallons) Mililiter {
	return Mililiter(g * 3785.41)
}

/*But those names would be a pain to write out all the time, and as we start adding
functions to convert between the other types, it becomes clear this isn’t
sustainable.*/