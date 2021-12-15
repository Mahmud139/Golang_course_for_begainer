package main

import "fmt"

func main(){
	//slide no 3
	fmt.Println(1)
	fmt.Println(1,2,3,4,5)
	letters := []string{"a"}
	letters = append(letters, "b")
	letters = append(letters, "c","d","f")
	fmt.Println(letters)

	//twoInts(1) //we can't pass just 1 arguments
	//twoInts(1,2,3) //we can't pass just 3 arguments
	twoInts(1,2) //we must pass just two arguments because of this function receive two parameter

	//slide no 5
	severalInts(1)
	severalInts(1,2,3,4,5)
	severalInts()

	//slide no 6
	severalString("a")
	severalString("a","b","c","d","e")
	severalString()


	//slide no 7
	mix(1, true, "a", "b")
	mix(2, false)

}
//slide no 3
func twoInts(first int, second int){
	fmt.Println(first, second)
}

//slide no 4
// func myFunc(param1 int, param2 ...int){
// 	//function code here
// }

//slide no 5
func severalInts(numbers ...int){
	fmt.Println(numbers)
}
//slide no 6
func severalString(name ...string){
	fmt.Println(name)
}

//slide no 7
func mix(num int, flag bool, names ...string){
	fmt.Println(num, flag, names)
}