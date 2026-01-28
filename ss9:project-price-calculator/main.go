package main

import (
	"fmt"

	"gmo.com/calculator/cmdmanager"
	"gmo.com/calculator/filemanager"
	"gmo.com/calculator/prices"
)

type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data any) error 
}

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		priceJob.Process()
	}
}
