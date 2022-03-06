// How to tell zero values apart from assigned value
package main

import "fmt"

func main(){
	//slide no 15
	// status("hasan")
	// status("iren")

	//slide no 16
	// counters := map[string]int{"a": 3, "b": 0}
	// var value int
	// var ok bool
	// value, ok = counters["a"]
	// fmt.Println(value, ok)
	// value, ok = counters["b"]
	// fmt.Println(value, ok)
	// value, ok = counters["c"]
	// fmt.Println(value, ok)

	//slide no 17
	// counters := map[string]int{"a": 3, "b": 0}
	// var ok bool
	// _, ok = counters["b"]
	// fmt.Println(ok)
	// _, ok = counters["c"]
	// fmt.Println(ok)

	//slide no 15
	status("hasan")
	status("iren")
}

// func status (name string) {
// 	grades := map[string]float64{"mahmud": 67.12, "hasan": 56.23,}
// 	grade := grades[name]
// 	if grade < 60 {
// 		fmt.Printf("%v is fail\n",name)
// 	}
// }

//slide no 18
func status (name string) {
	grades := map[string]float64{"mahmud": 67.12, "hasan": 56.23,}
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("No grade recorded for %s\n",name)
	}else if grade < 60 {
		fmt.Printf("%v is fail\n",name)
	}
}