package main

import (
	"bufio"
	"fmt"
	"github.com/mhaatha/go-cli-calculator/format"
	"github.com/mhaatha/go-cli-calculator/operations"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var firstNumber, secondNumber, result float64
	var operationName, yesOrNo string

	fmt.Print(format.App)

	for {
		fmt.Print(format.Operation)

		// Function to choose operation
		// Return string (example: "1", "2", etc.)
		operationNumber := operations.OperationNumbers(scanner)

		// Function to fill operationName variable
		operations.OperationCases(operationNumber, &operationName)
		// Error handling if operationName is empty
		if operationName == "" {
			fmt.Print(format.Invalid)
			time.Sleep(time.Second * 2)
			continue
		}

		// These two function below will parse the firstNumber and secondNumber variable
		// from string into float64
		operations.FirstNum(scanner, &firstNumber)
		operations.SecondNum(scanner, &secondNumber)

		// Function to calculate the result
		operations.Calculate(operationName, &result, &firstNumber, &secondNumber)

		fmt.Printf("\nThe result of the %v is: %.2f \n", operationName, result)

		// Function to ask the user to continue or not
		operations.IsContinue(scanner, &yesOrNo)

		if strings.ToLower(yesOrNo) == "no" {
			break
		}
	}
}
