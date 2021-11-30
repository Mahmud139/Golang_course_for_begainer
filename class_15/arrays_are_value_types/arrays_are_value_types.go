package main

import "fmt"

func main(){
	// name := [3]string{"mahmud","shihab","hasan"}
	// b := name
	// b[0] = "hasan"

	// fmt.Println(name)
	// fmt.Println(b)

	num := [5]int{5, 6, 7, 8, 8}
    fmt.Println("before passing to function ", num)
    changeLocal(num) //num is passed by value
    fmt.Println("after passing to function ", num)

}

func changeLocal(num [5]int) {  
    num[0] = 55
    fmt.Println("inside function ", num)
}