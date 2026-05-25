package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, t := range tokens {
		if t == "+" || t == "-" || t == "*" || t == "/" {
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			switch t {
			case "+":
				stack = append(stack, b+a)
			case "-":
				stack = append(stack, b-a)
			case "*":
				stack = append(stack, b*a)
			case "/":
				stack = append(stack, b/a)
			}
		} else {
			num, _ := strconv.Atoi(t)
			stack = append(stack, num)
		}
	}

	return stack[0]
}
