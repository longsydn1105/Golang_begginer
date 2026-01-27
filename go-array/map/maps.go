package main

import "fmt"

func main() {
	websites := make(map[string]string)

	websites["Google"] = "https://google.com"
	websites["Facebook"] = "https://facebook.com"

	fmt.Println(websites)
	fmt.Println(websites["Google"])

	delete(websites, "Facebook")
	fmt.Println(websites)
}
