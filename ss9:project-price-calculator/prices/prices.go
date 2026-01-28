package prices

import (
	"fmt"

	"gmo.com/calculator/conversion"
	"gmo.com/calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager        iomanager.IOManager `json:"-"`
	TaxRate          float64             `json:"tax_rate"`
	InputPrices      []float64           `json:"input_prices"`
	TaxIncludePrices map[string]string   `json:"tax_included_price"`
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool) {
	err := job.LoadData()
	if err != nil {
		//return
	}
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncluedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncluedPrice)
	}

	job.TaxIncludePrices = result
	job.IOManager.WriteResult(job)
	doneChan <- true
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
