package prices

import (
	"bufio"
	"fmt"
	"os"

	"github.com/benkigera/price_calculator/conversions"
)

type TaxIncludedPriceJob struct {
	TaxRate float64
	InputPrices  []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) ReadInputPrices() {
	file, err := os.Open("price.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		price, err := conversions.StringToFloat(line)
		if err != nil {
			if line != "" {
				fmt.Println("Error parsing price:", err)
			}
			continue
		}
		job.InputPrices = append(job.InputPrices, price)
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.ReadInputPrices()
	job.TaxIncludedPrices = make(map[string]float64)
	
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		job.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = taxIncludedPrice
	}

	fmt.Println("Price Calculations:")
	for originalPrice, taxIncludedPrice := range job.TaxIncludedPrices {
		fmt.Printf("Original: $%s -> With Tax: $%.2f\n", originalPrice, taxIncludedPrice)
	}
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}