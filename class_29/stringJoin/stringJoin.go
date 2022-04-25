package main

import (
	"fmt"
	"strings"
)

func main() {
	name := []string{"Mahmud", "Hasan"}
	res := strings.Join(name, " + ")
	fmt.Println(res)
}