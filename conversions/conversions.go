package conversions

import (
	"fmt"
	"strconv"
)

// StringToFloat converts a string to float64, returning an error if the conversion fails
// or if the string is empty
func StringToFloat(s string) (float64, error) {
	if s == "" {
		return 0, fmt.Errorf("empty string cannot be converted to float")
	}
	
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to convert string to float: %w", err)
	}
	
	return value, nil
}