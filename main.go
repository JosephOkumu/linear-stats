package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Function to calculate the mean of a slice of float64
func Mean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

// function to calculate linear regression
func linearRegression(inputData []float64) (float64, float64) {
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

// function to calculate pearson correlation
func pearsonCorrelation(inputData []float64) float64 {
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

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: go run main.go <filename>")
		return
	}
	data := os.Args[1]

	file, err := os.Open(data)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var numbers []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", err)
			continue
		}
		numbers = append(numbers, num)
	}

	m, b := linearRegression(numbers)
	pearsonCor := pearsonCorrelation(numbers)

	fmt.Printf("Linear regression line: y = %.6fx + %.6f\n", m, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", pearsonCor)

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}
