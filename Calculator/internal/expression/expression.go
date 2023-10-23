package expression

import (
	"fmt"
	"strings"

	"github.com/johncgriffin/overflow"
)

func ScanExpression(expr string) (Expression, error) {
	var (
		firstNumber, secondNumber int
		operator                  rune
		exp                       Expression
	)

	reader := strings.NewReader(expr)
	if _, err := fmt.Fscanf(reader, "%d %c %d\n", &firstNumber, &operator, &secondNumber); err != nil {
		return nil, err
	}
	switch operator {
	case '+':
		exp = &Addition{LeftOperand: firstNumber, RightOperand: secondNumber}
	case '-':
		exp = &Substraction{LeftOperand: firstNumber, RightOperand: secondNumber}
	case '*':
		exp = &Multiplication{LeftOperand: firstNumber, RightOperand: secondNumber}
	case '/':
		exp = &Division{LeftOperand: firstNumber, RightOperand: secondNumber}
	default:
		return nil, fmt.Errorf("operation %c is not provided", operator)
	}
	return exp, nil
}

type Expression interface {
	Evaluate() (int, error)
}

type Addition struct {
	LeftOperand  int
	RightOperand int
}

func (add *Addition) Evaluate() (int, error) {
	if result, ok := overflow.Add(add.LeftOperand, add.RightOperand); ok {
		return result, nil
	}
	return 0, fmt.Errorf("addition wasn't complete with operators: %d, %d", add.LeftOperand, add.RightOperand)
}

type Multiplication struct {
	LeftOperand  int
	RightOperand int
}

func (add *Multiplication) Evaluate() (int, error) {
	if result, ok := overflow.Mul(add.LeftOperand, add.RightOperand); ok {
		return result, nil
	}
	return 0, fmt.Errorf("multiplication wasn't complete with operators: %d, %d", add.LeftOperand, add.RightOperand)
}

type Division struct {
	LeftOperand  int
	RightOperand int
}

func (add *Division) Evaluate() (int, error) {
	if result, ok := overflow.Div(add.LeftOperand, add.RightOperand); ok {
		return result, nil
	}
	return 0, fmt.Errorf("division wasn't complete with operators: %d, %d", add.LeftOperand, add.RightOperand)
}

type Substraction struct {
	LeftOperand  int
	RightOperand int
}

func (add *Substraction) Evaluate() (int, error) {
	if result, ok := overflow.Sub(add.LeftOperand, add.RightOperand); ok {
		return result, nil
	}
	return 0, fmt.Errorf("substraction wasn't complete with operators: %d, %d", add.LeftOperand, add.RightOperand)
}
