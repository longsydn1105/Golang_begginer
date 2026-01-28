package prices

import (
	"fmt"

	"gmo.com/calculator/conversion"
	"gmo.com/calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager        filemanager.FileManager `json:"-"`
	TaxRate          float64                 `json:"tax_rate"`
	InputPrices      []float64               `json:"input_prices"`
	TaxIncludePrices map[string]string       `json:"tax_included_price"`
}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncluedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncluedPrice)
	}

	job.TaxIncludePrices = result
	job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
