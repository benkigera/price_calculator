package main

import (
	"github.com/benkigera/price_calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1}

	for _, taxRate := range taxRates {
	pricejob :=	prices.NewTaxIncludedPriceJob(taxRate)

	pricejob.Process()
	}

	
}
