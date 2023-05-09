//maps are reference type same as slice
package main

import "fmt"

func main(){
	employeeSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,     
        "mike": 9000,
    }
	// fmt.Println("Original employee salary", employeeSalary)
    // modified := employeeSalary
    // modified["mike"] = 18000
    // fmt.Println("Employee salary changed", employeeSalary)


    fmt.Println("before calling change(): ", employeeSalary)
    change(employeeSalary)
    fmt.Println("after calling change():", employeeSalary)
}

func change(m map[string]int) {
    m["jamie"] = 10
    fmt.Println("inside the change(): ", m)
}