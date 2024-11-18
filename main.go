package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mhaatha/go-cli-calculator/format"
	"github.com/mhaatha/go-cli-calculator/operations"
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

		operations.Result(operationName, result)

		// Function to ask the user to continue or not
		operations.IsContinue(scanner, &yesOrNo)

		if strings.ToLower(yesOrNo) == "no" {
			break
		} else if strings.ToLower(yesOrNo) == "yes" {
			continue
		} else {
			fmt.Print(format.Invalid)
			time.Sleep(time.Second * 2)
			continue
		}
	}
}
