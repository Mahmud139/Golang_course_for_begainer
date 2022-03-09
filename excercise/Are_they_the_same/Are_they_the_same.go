package main

import (
	"fmt"
	"sort"
)

//a = [121, 144, 19, 161, 19, 144, 19, 11]  
//b = [121, 14641, 20736, 361, 25921, 361, 20736, 361]

func Comp(array1 []int, array2 []int) bool {
	if array1 == nil || array2 == nil{
		return false
	}
    for i,v:=range array1{
		array1[i] = v*v
	}

	sort.Ints(array1)
	sort.Ints(array2)

	return Equal(array1,array2)
}

func Equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func main(){
	a := []int{121, 144, 19, 161, 19, 144, 19, 11} 
	b := []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}

	fmt.Println(Comp(a,b))

}