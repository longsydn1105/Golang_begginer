package main

import "fmt"

func main() {
	sum := sumup(1, 4, 14, 44, 444)
	fmt.Println(sum)
}

func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
