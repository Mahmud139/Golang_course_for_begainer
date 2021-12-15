package main
import "fmt"

func main() {
	average([]int{89, 90, 95})
	average([]int{56, 67, 45, 90, 109})
	average([]int{38, 56, 98})
	average([]int{})
}

func average(numbers []int) {
	if len(numbers) == 0 {
		return
	}

	sum := 0
	for _, number := range numbers {
		sum += number
	}
	
	fmt.Println(sum / len(numbers))
}
