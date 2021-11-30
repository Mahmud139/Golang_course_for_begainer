package main 

import "fmt"

func main(){
	//name := "Shihab"
	array := [5]int{1,2,3,4,5}
	fmt.Println(array[2])
	// 0 1234 5
	//make sure that the string does not contain outside of ASCII character list
	// 1 i<6
	// for i:= 1; i< len(name)-1; i++ {
	// 	fmt.Printf("%c",name[i])
	// }
	// var i int = 12
	// a := float64(i)

	// r := removeCharacter("shihab")
	// fmt.Printf(r)
	// r := repeat(6, "I")
	// fmt.Printf(r)
}

// func removeCharacter(s string) string {
// 	var res string
// 	for i:= 1; i< len(s)-1; i++ {
// 		//fmt.Println(fmt.Sprintf("%c",s[i]))
// 		//res = res + fmt.Sprintf("%c",s[i])
// 		res = res + string(s[i])
// 	}
// 	return res
// }


// func repeat(n int, s string) string {
// 	var res string
// 	for i:=0;i<n; i++ {
// 		res = res + s
// 	}
// 	return res
// }

