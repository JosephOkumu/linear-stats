package calcs

import "math"

// function to calculate pearson correlation
func PearsonCorrelation(inputData []float64) float64 {
	x := make([]float64, len(inputData))

	// Generate x axis values
	for i := range inputData {
		x[i] = float64(i)
	}

	// Calculate means
	meanX := Mean(x)
	meanY := Mean(inputData)

	var sumXY, squareX, squareY float64
	for i, yi := range inputData {
		xi := float64(i)
		sumXY += (xi - meanX) * (yi - meanY)
		squareX += (xi - meanX) * (xi - meanX)
		squareY += (yi - meanY) * (yi - meanY)
	}

	// Calculate coefficient
	rootX := math.Sqrt(squareX)
	rootY := math.Sqrt(squareY)
	rootXy := rootX * rootY
	if rootXy == 0 {
		return 0
	}

	correlation := sumXY / rootXy
	return correlation
}