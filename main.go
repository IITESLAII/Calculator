package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var exitLine []string
var stack []string

func Priority(operator string) int {
	switch operator {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "(":
		return 0
	}
	return -1
}

func GetStack(stack []string) string {
	if len(stack) == 0 {
		return ""
	}
	return stack[len(stack)-1]
}

func GetStackWithInt(stack []string, i int) string {
	if len(stack) == 0 {
		return ""
	}
	return stack[len(stack)-i]
}

func DeleteStack(s *[]string) {
	if len(*s) > 0 {
		*s = (*s)[:len(*s)-1]
	}
}

func ToStack(s *[]string, operator string) {
	if operator == "(" {
		*s = append(*s, operator)
		return
	}
	if operator == ")" {
		for len(*s) > 0 && GetStack(*s) != "(" {
			exitLine = append(exitLine, GetStack(*s))
			DeleteStack(s)
		}
		if len(*s) == 0 {
			exitLine = []string{}
			return
		}
		DeleteStack(s)
		return
	}

	for len(*s) > 0 && Priority(GetStack(*s)) >= Priority(operator) {
		exitLine = append(exitLine, GetStack(*s))
		DeleteStack(s)
	}
	*s = append(*s, operator)
}

func ToPostfix(expression string) ([]string, error) {
	re := regexp.MustCompile(`(\d+\.?\d*|[+\-*/()])`)
	sliceExpression := re.FindAllString(expression, -1)

	if len(sliceExpression) == 0 {
		return nil, fmt.Errorf("некорректное выражение")
	}

	lastWasOperator := true
	openParentheses := 0

	for _, s := range sliceExpression {
		if s == "+" || s == "-" || s == "*" || s == "/" {
			if lastWasOperator {
				return nil, fmt.Errorf("некорректное выражение: оператор не может следовать за оператором")
			}
			lastWasOperator = true
			ToStack(&stack, s)
		} else if s == "(" {
			lastWasOperator = true
			openParentheses++
			ToStack(&stack, s)
		} else if s == ")" {
			lastWasOperator = false
			ToStack(&stack, s)
			if openParentheses == 0 {
				return nil, fmt.Errorf("некорректное выражение: лишняя закрывающая скобка")
			}
			openParentheses--
		} else {
			lastWasOperator = false
			exitLine = append(exitLine, s)
		}
	}

	if openParentheses > 0 {
		return nil, fmt.Errorf("некорректное выражение: отсутствует закрывающая скобка")
	}

	for len(stack) != 0 {
		if lastWasOperator {
			return nil, fmt.Errorf("некорректное выражение: выражение заканчивается на оператор")
		}
		exitLine = append(exitLine, GetStack(stack))
		DeleteStack(&stack)
	}

	return exitLine, nil
}

func AddToStack(st *[]string, f float64) {
	stringNumber := fmt.Sprintf("%f", f)
	*st = append(*st, stringNumber)
}

func Calculate(postfixExp []string) ([]string, error) {
	st := make([]string, 0)
	for _, s := range postfixExp {
		switch s {
		case "+", "-", "/", "*":
			if len(st) < 2 {
				return nil, fmt.Errorf("некорректное выражение: недостаточно операндов")
			}
			x, _ := strconv.ParseFloat(GetStackWithInt(st, 2), 64)
			y, _ := strconv.ParseFloat(GetStackWithInt(st, 1), 64)
			DeleteStack(&st)
			DeleteStack(&st)

			var result float64
			switch s {
			case "+":
				result = x + y
			case "-":
				result = x - y
			case "/":
				if y == 0 {
					return nil, fmt.Errorf("ошибка: деление на ноль")
				}
				result = x / y
			case "*":
				result = x * y
			}
			AddToStack(&st, result)
		default:
			st = append(st, s)
		}
	}

	return st, nil
}

func Calc(expression string) (float64, error) {
	exitLine = []string{}
	stack = []string{}

	a := strings.ReplaceAll(expression, " ", "")
	postfixExpression, err := ToPostfix(a)
	if err != nil {
		return 0, err
	}
	resultStack, err := Calculate(postfixExpression)
	if err != nil {
		return 0, err
	}

	if len(resultStack) > 0 {
		finalResult, _ := strconv.ParseFloat(resultStack[0], 64)
		fmt.Println(finalResult)
		return finalResult, nil
	}
	return 0, nil
}

func main() {
	for true {
		a := ""
		fmt.Scan(&a)
		Calc(a)
	}
}
