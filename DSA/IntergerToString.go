package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(IntToString(1994))
}

func IntToString(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var sb strings.Builder

	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			num -= values[i]
			sb.WriteString(symbols[i])
		}
	}
	return sb.String()
}
