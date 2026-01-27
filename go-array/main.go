package main

import "fmt"

type floatmap map[string]float64

func (fm floatmap) output() {
	fmt.Println(fm)
}

func main() {
	// userNames := make([]string, 2)

	// userNames = append(userNames, "Max")
	// userNames = append(userNames, "Manuel")

	// fmt.Println(userNames)

	courseRatings := make(floatmap, 2)

	courseRatings["go"] = 4.7
	courseRatings["NodeJS"] = 4.1

	//fmt.Println(courseRatings)

	courseRatings.output()

	for index, value := range courseRatings {
		fmt.Println("Index: ", index)
		fmt.Println("Values: ", value)
	}
}
