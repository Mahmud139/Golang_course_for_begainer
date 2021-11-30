package main

import "fmt"

func main(){
	//slide no 36
	twoDSlice1 := make([][]int, 2)
	for i := range twoDSlice1 {
		twoDSlice1[i] = make([]int, 3)
	}
	// twoDSlice1[0] = make([]int, 3)
	// twoDSlice1[1] = make([]int, 3)
	// twoDSlice1[0] = []int{1, 2, 3}
	// twoDSlice1[1] = []int{4, 5, 6}

	fmt.Printf("Number of rows in slice: %d\n", len(twoDSlice1))
    fmt.Printf("Number of columns in slice: %d\n", len(twoDSlice1[0]))
    fmt.Printf("Total number of elements in slice: %d\n", len(twoDSlice1)*len(twoDSlice1[0]))
	fmt.Printf("First Slice:%v\n", twoDSlice1[0])
	fmt.Printf("Second Slice:%v\n", twoDSlice1[1])



	//slide no 37
	// twoDSlice := make([][]int, 2)
    // twoDSlice[0] = []int{1, 2, 3}
    // twoDSlice[1] = []int{4, 5}
	// twoDSlice[0] = make([]int, 3)
    // twoDSlice[1] = make([]int, 2)

    // fmt.Printf("Number of rows in slice: %d\n", len(twoDSlice))
    // fmt.Printf("Len of first row: %d\n", len(twoDSlice[0]))
    // fmt.Printf("Len of second row: %d\n", len(twoDSlice[1]))
	// fmt.Printf("First Slice:%v\n", twoDSlice[0])
	// fmt.Printf("Second Slice:%v\n", twoDSlice[1])

	// fmt.Println("Traversing slice")
    // for _, row := range twoDSlice {
    //     for _, val := range row {
    //         fmt.Println(val)
    //     }
    // }
}