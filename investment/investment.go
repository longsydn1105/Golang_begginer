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

    formattedFV := fmt.Sprintf("Future Value     : %2.2f\n", futureValue)
    formattedRFV := fmt.Sprintf("Future Real Value: %2.2f\n", realFutureVale)

    fmt.Print(formattedFV, formattedRFV)
    // fmt.Printf("Future Value: %v \n",futureValue)
    // fmt.Printf("FutureRealValue: %+v \n",realFutureVale)
}