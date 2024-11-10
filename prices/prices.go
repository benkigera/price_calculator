package prices

import (
	"fmt"

	"github.com/benkigera/price_calculator/filemanager"
)

type PriceResult struct {
	Original float64
	WithTax  float64
}

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices []PriceResult
}

func (job *TaxIncludedPriceJob) ReadInputPrices() {
	prices, err := filemanager.ReadPricesFromFile("price.txt")
	if err != nil {
		fmt.Println("Error reading prices:", err)
		return
	}
	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.ReadInputPrices()
	job.TaxIncludedPrices = make([]PriceResult, 0, len(job.InputPrices))
	
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		job.TaxIncludedPrices = append(job.TaxIncludedPrices, PriceResult{
			Original: price,
			WithTax:  taxIncludedPrice,
		})
	}

	fmt.Println("Price Calculations:")
	for _, result := range job.TaxIncludedPrices {
		fmt.Printf("Original: $%.2f -> With Tax: $%.2f\n", result.Original, result.WithTax)
	}
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}