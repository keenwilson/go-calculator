package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Create a single reader object
var reader = bufio.NewReader(os.Stdin)

func main() {
	value1 := getInput("Value 1")
	value2 := getInput("Value 2")

	var result float64

	switch operation := getOperation(); operation {
	case "+":
		result = addValues(value1, value2)
	case "-":
		result = subtractValues(value1, value2)
	case "*":
		result = multiplyValues(value1, value2)
	case "/":
		result = divideValues(value1, value2)
	default:
		panic("invalid operation")
	}
	result = math.Round(result*100) / 100
	fmt.Printf("The result is %v\n\n", result)

}

func getInput(propmt string) float64 {
	fmt.Printf("%v ", propmt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message := fmt.Sprintf("%v must be a number", propmt)
		panic(message)
	}
	return value
}

func getOperation() string {
	fmt.Print("Select an operation (+ - * /): ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace((op))
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
