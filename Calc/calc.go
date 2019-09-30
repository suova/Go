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

func parseInfix(equation []string) (rpnString string) {
	stack := make(Stack, 0)
	for _, elem := range equation {
		switch elem {
		case "(":
			stack.push(elem)
		case ")":

			for {
				operand, err := stack.pop()
				if err != nil {
					return ""
				}
				if operand == "(" {
					break // отрасываем "("
				}
				rpnString += " " + operand
			}
		default:
			if elem1, isOperator := operations[elem]; isOperator {
				//если elem1 оператор, заходим сюда
				for len(stack) > 0 {
					topElem := stack[len(stack)-1]
					if elem2, isOperator := operations[topElem]; !isOperator || elem1.precedence > elem2.precedence ||
						elem1.precedence == elem2.precedence {
						break
					}

					stack = stack[:len(stack)-1]
					rpnString += " " + topElem
				}
				stack = append(stack, elem)
			} else {
				if rpnString > "" {
					rpnString += " "
				}
				rpnString += elem
			}
		}
	}

	for len(stack) > 0 {
		elem, restStack := stack[len(stack)-1], stack[:len(stack)-1]
		rpnString += " " + elem
		stack = restStack
	}
	return
}
func calc(equation []string) (string, error) {
	stack := make(Stack, 0)

	for _, elem := range equation {
		i, err := strconv.Atoi(elem)
		if err == nil {
			stack.push(strconv.Itoa(i))
		} else {
			switch elem {
			case "+", "-", "*", "/":

				a, err := stack.pop()
				if err != nil {
					return "", err
				}
				b, err := stack.pop()
				if err != nil {
					return "", err
				}

				elem1, _ := strconv.Atoi(b)
				elem2, _ := strconv.Atoi(a)

				switch elem {
				case "+":
					fmt.Println(elem1 + elem2)
					stack.push(strconv.Itoa(elem1 + elem2))

				case "-":
					stack.push(strconv.Itoa(elem1 - elem2))

				case "*":
					stack.push(strconv.Itoa(elem1 * elem2))

				case "/":
					if elem2 != 0 {
						stack.push(strconv.Itoa(elem1 / elem2))

					} else {
						return "", errors.New(fmt.Sprintf("division by zero"))
					}
				default:
					return "", errors.New(fmt.Sprintf("This symbol is not used in calc '%s'", elem))
				}
			default:
				return "", errors.New(fmt.Sprintf("This symbol is not used in calc '%s'", elem))
			}
		}
	}

	result, err := stack.pop()
	if err != nil {
		return "", err
	}
	return result, nil
}

type Stack []string

func (stack *Stack) pop() (string, error) {
	l := len(*stack)
	if l == 0 {
		return "", errors.New("len of stack is null")
	}
	elem := (*stack)[l-1]
	*stack = (*stack)[:l-1]
	return elem, nil
}
func (stack *Stack) push(elem string) {
	*stack = append(*stack, elem)
}

func main() {
	expression := strings.Split((os.Args[1:])[0], "")
	if len(expression) == 0 {
		fmt.Println("length is null")
	}

	equation := parseInfix(expression)
	result, err := calc(strings.Split(equation, " "))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
