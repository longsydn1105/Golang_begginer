package main

import (
	"fmt"
	"math"
)

func main() {
    const inflationRate float64 = 2.5
    var investmentAmount float64
    var Years float64
    expectedReturnRate := 5.5

    fmt.Print("InvestmentAmount: ")
    fmt.Scan(&investmentAmount)

    fmt.Print("Years: ")
    fmt.Scan(&Years)

    fmt.Print("ExpectedReturnRate: ")
    fmt.Scan(&expectedReturnRate)

    futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, Years)
    realFutureVale := futureValue / math.Pow(1+inflationRate/100,  Years)
    fmt.Println(futureValue)
    fmt.Println(realFutureVale)
}