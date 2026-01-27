package main

import "fmt"

type trasnsformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	// doubled := transformNumbers(&numbers, double)
	// fmt.Println(double)
	transformerFn1 := getTransformFn(&numbers)
	transformerFn2 := getTransformFn(&moreNumbers)

	transformedNumbers1 := transformNumbers(&numbers, transformerFn1)
	transformedNumbers2 := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(transformedNumbers1)
	fmt.Println(transformedNumbers2)
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
func getTransformFn(num *[]int) trasnsformFn {
	if (*num)[0] == 1 {
		return double
	} else {
		return triple
	}
}
