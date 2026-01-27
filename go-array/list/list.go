package list

import "fmt"

// type Product struct {
// 	title string
// 	id    string
// 	price float64
// }

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices = append(prices, 5.99, 12.99, 4.44)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{10.99, 8.99, 44.44, 9.2}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// func main() {
// 	// 1)
// 	hobby := [3]string{"code", "gym", "sleep"}

// 	//2)
// 	fmt.Println(hobby[0])
// 	fmt.Println(hobby[1:])

// 	// 3)
// 	hobby2 := hobby[:2]
// 	//hobby3 := hobby[0:2]

// 	//4)
// 	hobby2 = hobby2[1:3]
// 	fmt.Println(hobby2)

// 	//5)
// 	goals := []string{"code", "pro-golang", "love-golang"}
// 	fmt.Println(goals)

// 	// 6)
// 	goals[1] = "senior-golang"
// 	fmt.Println(goals)

// 	myGoals := append(goals, "golang-in-heart")
// 	fmt.Println(myGoals)
// 	fmt.Println(goals)

// 	//bai 7
// 	pro1 := Product{
// 		"Cay vot",
// 		"sp1",
// 		50.0,
// 	}

// 	pro2 := Product{
// 		"qua cau",
// 		"sp2",
// 		10.0,
// 	}
// 	listproducts := []Product{pro1, pro2}
// 	fmt.Println(listproducts)

// 	pro3 := Product{
// 		"cai nit",
// 		"sp3",
// 		0.0,
// 	}
// 	listproducts = append(listproducts, pro3)
// 	fmt.Println(listproducts)

// }

// func main() {
// 	prices := []float64{10.99, 8.99}
// 	fmt.Println(prices[1])
// 	prices[1] = 9.99
// 	updatePrices := append(prices,5.99)
// 	fmt.Println(updatePrices, prices)
// }

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	productNames[2] = "A Carpet"

// 	prices := [4]float64{10.99, 9.99, 46.99, 20.0}

// 	fmt.Println(prices[1])
// 	fmt.Println(productNames[2])

// 	featuresPrices := prices[:1]
// 	featuresPrices[0] = 199.99
// 	fmt.Println(featuresPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(featuresPrices), cap(featuresPrices))
// }
