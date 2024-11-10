package filemanager

import (
	"bufio"
	"os"

	"github.com/benkigera/price_calculator/conversions"
)

func ReadPricesFromFile(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var prices []float64
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()
		price, err := conversions.StringToFloat(line)
		if err != nil {
			// Skip empty lines silently
			if line != "" {
				// Return error for non-empty invalid lines
				return nil, err
			}
			continue
		}
		prices = append(prices, price)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return prices, nil
}