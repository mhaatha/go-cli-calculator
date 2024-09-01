package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var firstNumber, secondNumber, result float64
	var operationName, yesOrNo string

	fmt.Println("===============================")
	fmt.Println("       🖩CLI Calculator🖩        ")
	fmt.Println("===============================")

	for {
		fmt.Println("\nSelect an operation:")
		fmt.Println("1. Addition (+)")
		fmt.Println("2. Subtraction (-)")
		fmt.Println("3. Multiplication (*)")
		fmt.Println("4. Division (/)")

		fmt.Print("\nEnter the number of the operation you want to perform: ")
		scanner.Scan()
		operation := scanner.Text()

		switch operation {
		case "1":
			operationName = "addition"
		case "2":
			operationName = "subtraction"
		case "3":
			operationName = "multiplication"
		case "4":
			operationName = "division"
		default:
			fmt.Println("\n===============================")
			fmt.Println("     ⚠️ Invalid Operation⚠️     ")
			fmt.Println("===============================")
			time.Sleep(time.Second * 2)
			continue
		}

		fmt.Print("\nEnter the first number: ")
		scanner.Scan()
		inputOne := scanner.Text()

		fmt.Print("\nEnter the second number: ")
		scanner.Scan()
		inputTwo := scanner.Text()

		if s, err := strconv.ParseFloat(inputOne, 64); err == nil {
			firstNumber = s
		}

		if s, err := strconv.ParseFloat(inputTwo, 64); err == nil {
			secondNumber = s
		}

		switch operationName {
		case "addition":
			result = firstNumber + secondNumber
		case "subtraction":
			result = firstNumber - secondNumber
		case "multiplication":
			result = firstNumber * secondNumber
		default:
			result = firstNumber / secondNumber
		}

		fmt.Printf("\nThe result of the %v is: %.2f \n", operationName, result)

		fmt.Print("\nDo you want to perform another operation? (yes/no) ")
		scanner.Scan()
		yesOrNo = scanner.Text()
		if strings.ToLower(yesOrNo) == "no" {
			break
		}
	}
}
