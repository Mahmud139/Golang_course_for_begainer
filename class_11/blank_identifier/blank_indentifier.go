package main

import "fmt"

func main(){
	//area := findAreaPerimeter(10.8, 5.8)
	area, _ := findAreaPerimeter(10.8, 5.8)
	fmt.Printf("Area is %0.2f", area)

}

func findAreaPerimeter(length, width float64) (float64, float64) {
	area := length * width
	perimeter := (length + width) * 2
	return area, perimeter
}