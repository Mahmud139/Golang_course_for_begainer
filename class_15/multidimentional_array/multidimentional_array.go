package main 

import "fmt"

func main(){
	// a := [3][2]string{
    //     {"lion", "tiger"},
    //     {"cat", "dog"},
    //     {"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
    // }

    // fmt.Println(a)

	// for _, v1 := range a {
    //     // for _, v2 := range v1 {
    //     //     fmt.Printf("%s ", v2)
    //     // }
    //     // fmt.Printf("\n")
    //     fmt.Println(v1)
    // }


	// var b [3][2]string
    // b[0][0] = "apple"
    // b[0][1] = "samsung"
    // b[1][0] = "microsoft"
    // b[1][1] = "google"
    // b[2][0] = "AT&T"
    // b[2][1] = "T-Mobile"
	
	// for _, v1 := range b {
    //     for _, v2 := range v1 {
    //         fmt.Printf("%s ", v2)
    //     }
    //     fmt.Printf("\n")
    // }
	
	
	//set array's elements using loop
	var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

	arr:= [3][3]string{
		{"C#", "C", "Python"}, 
		{"Java", "Scala", "Perl"},
		{"C++", "Go", "HTML"},
	}
	fmt.Println(arr)
	
}