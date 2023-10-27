package main

import (
	"bufio"
	"fmt"
	"go/course/calculator/internal/expression"
	"io"
	"os"
)

func main() {
	availableOperations := []rune{'+', '-', '/', '*'}
	fmt.Println("This is Calculator CUI. Input EOF to stop the programm")
	fmt.Printf("Available operations: %c\n\n", availableOperations)
	stdin := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Input <first operand> <operator> <second operand>: ")
		var (
			line string
			err  error
		)
		if line, err = stdin.ReadString('\n'); err != nil {
			if err == io.EOF {
				fmt.Println("Exiting program")
				return
			}
			fmt.Printf("Error: %s\n\n", err)
			continue
		}
		expr, err := expression.ScanExpression(line)
		if err != nil {
			fmt.Printf("Error: %s\n\n", err)
			continue
		}
		result, err := expr.Evaluate()
		if err != nil {
			fmt.Printf("Error: %s\n\n", err)
		} else {
			fmt.Printf("Result: %d\n\n", result)
		}
	}
}
