package main

import "fmt" 

func main(){
	//repeatLine("Mahmud", 3)
	repeatLine("Shihab", 5)
	calculateBill(30,6)
}

func repeatLine (line string, times int) {
	for i:= 0; i < times; i++ {
		fmt.Println(line)
	}
}

/*If consecutive parameters are of the same type, we can avoid writing the 
type each time and it is enough to be written once at the end.
ie price int, no int can be written as price, no int.*/

func calculateBill(price, no int) {  
    var totalPrice = price * no
    fmt.Println(totalPrice)
}
