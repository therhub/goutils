package mathutil

import "math"

// save one small number
func ConvertFloat64ToDecimal(f float64) float64 {
	return float64(math.Floor((f+0.05)*10) / 10)
}
