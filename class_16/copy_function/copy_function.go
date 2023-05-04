package main

import "fmt"

func main(){
	src := []int{1, 2, 3, 4, 5}
    dst := make([]int, 5)

    /*Destination can't be nill or empty slice*/
    // var dst []int //--nill slice
    // dst := []int{} //--empty slice

    /*if destination slice already had elements then after copy, all previous value will be replaced by source elements*/
    // dst := []int{200}


    numberOfElementsCopied := copy(dst, src)
    fmt.Printf("Number Of Elements Copied: %d\n", numberOfElementsCopied)
    fmt.Printf("dst: %v\n", dst)
    fmt.Printf("src: %v\n", src)

    //After changing numbers2
    dst[0] = 10
    fmt.Println("\nAfter changing dst")
    fmt.Printf("dst: %v\n", dst)
    fmt.Printf("src: %v\n", src)

    /* 
    1. If the length of src is greater than the length of dst, then the number of elements copied is the length of dst.
    2. If the length of dst is greater than the length of src, then the number of elements copied is the length of src
    */
}