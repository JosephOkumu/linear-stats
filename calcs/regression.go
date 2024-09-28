package calcs

// function to calculate linear regression
func LinearRegression(inputData []float64) (float64, float64) {
	length := float64(len(inputData))
	var addY float64

	for _, y := range inputData {
		addY += y
	}

	
	meanX := (length - 1) / 2 
	meanY := Mean(inputData)

	
	numerator := 0.0
	denominator := 0.0

	for i := 0; i < int(length); i++ {
		numerator += (float64(i) - meanX) * (inputData[i] - meanY)
		denominator += (float64(i) - meanX) * (float64(i) - meanX)
	}

	m := numerator / denominator // Slope
	b := meanY - m*meanX        // Intercept

	return m, b
}