package main

import (
	"bufio"
	"fmt"
	"linear/calcs"
	"os"
	"strconv"
)

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

	m, b := calcs.LinearRegression(numbers)
	pearsonCor := calcs.PearsonCorrelation(numbers)

	fmt.Printf("Linear regression line: y = %.6fx + %.6f\n", m, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", pearsonCor)

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}
