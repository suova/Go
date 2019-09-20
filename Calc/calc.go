package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var operations = map[string]struct {
	precedence int
}{
	"*": {3},
	"/": {3},
	"+": {2},
	"-": {2},
}

func pushString(stack []string, elem string) []string {
	return append(stack, elem)
}
func popString(stack []string) (string, []string) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}

func parseInfix(equation []string) (rpnString string) {
	var stack []string
	for _, elem := range equation {
		switch elem {
		case "(":
			stack = pushString(stack, elem)
		case ")":
			var op string
			for {
				op, stack = popString(stack)
				if op == "(" {
					break // отрасываем "("
				}
				rpnString += " " + op
			}
		default:
			if o1, isOp := operations[elem]; isOp {
				for len(stack) > 0 {
					op, _ := popString(stack)
					if o2, isOp := operations[op]; !isOp || o1.precedence > o2.precedence ||
						o1.precedence == o2.precedence {
						break
					}
					_, stack = popString(stack)
					rpnString += " " + op
				}
				stack = pushString(stack, elem)
			} else {
				if rpnString > "" {
					rpnString += " "
				}
				rpnString += elem
			}
		}
	}

	for len(stack) > 0 {
		elem, restStack := popString(stack)
		rpnString += " " + elem
		stack = restStack
	}
	return
}
func calc(equation []string) (float64, error) {
	stack := make([]int, 0)

	for _, elem := range equation {
		i, err := strconv.Atoi(elem)

		if err == nil {
			stack = push(stack, i)
		} else {
			switch elem {
			case "+", "-", "*", "/":
				stack, err = action(stack, elem)
				if err != nil {
					return 0, err
				}
				break
			case " ":
				break
			default:
				return 0, errors.New(fmt.Sprintf("This symbol is not used in calc '%s'", elem))
			}
		}
	}

	result, _ := pop(stack)
	return float64(result), nil
}

func action(stack []int, operator string) ([]int, error) {
	elem1, stack := pop(stack)
	elem2, stack := pop(stack)
	switch operator {
	case "+":
		return append(stack, elem1+elem2), nil
	case "-":
		return append(stack, elem2-elem1), nil
	case "*":
		return append(stack, elem1*elem2), nil
	case "/":
		if elem1 != 0 {
			return append(stack, elem2/elem1), nil

		} else {
			return nil, errors.New(fmt.Sprintf("division by zero"))
		}
	default:
		return nil, errors.New(fmt.Sprintf("This symbol is not used in calc '%s'", operator))
	}
}

func pop(stack []int) (int, []int) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}
func push(stack []int, elem int) []int {
	return append(stack, elem)
}

func main() {
	s := strings.Split((os.Args[1:])[0], "")

	expression := ""
	for i, c := range s {
		if c == "(" && i == 0 {
			expression += c
			expression += " "
			continue
		} else if c == "*" || c == "/" || c == "+" || c == "-" || c == ")" || c == "(" {
			expression += " "
			expression += c
			expression += " "
		} else {
			expression += c
		}
	}

	equation := parseInfix(strings.Fields(expression))
	result, err := calc(strings.Split(equation, " "))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
