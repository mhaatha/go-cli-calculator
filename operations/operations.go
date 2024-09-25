package operations

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
)

func OperationNumbers(scanner *bufio.Scanner) string {
	fmt.Print("\nEnter the number of the operation you want to perform: ")
	scanner.Scan()

	return scanner.Text()
}

func OperationCases(operationNumber string, operationName *string) {
	switch operationNumber {
	case "1":
		*operationName = "addition"
	case "2":
		*operationName = "subtraction"
	case "3":
		*operationName = "multiplication"
	case "4":
		*operationName = "division"
	default:
		*operationName = ""
	}
}

func FirstNum(scanner *bufio.Scanner, firstNumber *float64) {
	fmt.Print("\nEnter the first number: ")
	scanner.Scan()

	inputOne := scanner.Text()

	if s, err := strconv.ParseFloat(inputOne, 64); err == nil {
		*firstNumber = s
	} else {
		log.Fatal(err)
	}
}

func SecondNum(scanner *bufio.Scanner, secondNumber *float64) {
	fmt.Print("\nEnter the second number: ")
	scanner.Scan()

	inputTwo := scanner.Text()

	if s, err := strconv.ParseFloat(inputTwo, 64); err == nil {
		*secondNumber = s
	} else {
		log.Fatal(err)
	}
}

func Calculate(operationName string, result *float64, firstNumber, secondNumber *float64) {
	switch operationName {
	case "addition":
		*result = *firstNumber + *secondNumber
	case "subtraction":
		*result = *firstNumber - *secondNumber
	case "multiplication":
		*result = *firstNumber * *secondNumber
	default:
		*result = *firstNumber / *secondNumber
	}
}

func IsContinue(scanner *bufio.Scanner, yesOrNo *string) {
	fmt.Print("\nDo you want to perform another operation? (yes/no) ")
	scanner.Scan()
	*yesOrNo = scanner.Text()
}
