package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	pentale := createTransformer(5)

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	pentaleNumbers := transformNumbers(&numbers, pentale)

	fmt.Println(transformed)
	fmt.Println(pentaleNumbers)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dbNumbers := []int{}

	for _, value := range *numbers {
		dbNumbers = append(dbNumbers, transform(value))
	}

	return dbNumbers
}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}

// func getTransformFn(num *[]int) transformFn {
// 	if (*num)[0] == 1 {
// 		return double
// 	} else {
// 		return triple
// 	}
// }

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
