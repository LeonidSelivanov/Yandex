package main

import (
	"errors"
	"strconv"
)

var errFoo = errors.New("error!")

func Calc(expression string) (float64, error) {
	if len(expression) == 0 {
		return 0, errFoo
	}
	var tokens []string
	for _, char := range expression {
		tokens = append(tokens, string(char))
	}
	rpnTokens, err := convertToRPN(tokens)
	if err != nil {
		return 0, err
	}
	result, err := evaluateRPN(rpnTokens)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func convertToRPN(tokens []string) ([]string, error) {
	var output []string
	var stack []string
	operators := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}
	for _, token := range tokens {
		if _, err := strconv.ParseFloat(token, 64); err == nil {
			output = append(output, token)
		} else if token == "(" {
			stack = append(stack, token)
		} else if token == ")" {
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil, errFoo
			}
			stack = stack[:len(stack)-1]
		} else if _, ok := operators[token]; ok {
			for len(stack) > 0 {
				if stack[len(stack)-1] == "(" || operators[stack[len(stack)-1]] < operators[token] {
					break
				}
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		} else {
			return nil, errFoo
		}
	}
	for len(stack) > 0 {
		if stack[len(stack)-1] == "(" {
			return nil, errFoo
		}
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return output, nil
}

func evaluateRPN(tokens []string) (float64, error) {
	var stack []float64
	for _, token := range tokens {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, errFoo
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var result float64
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, errFoo
				}
				result = a / b
			default:
				return 0, errFoo
			}
			stack = append(stack, result)
		}
	}
	if len(stack) != 1 {
		return 0, errFoo
	}
	return stack[0], nil
}
