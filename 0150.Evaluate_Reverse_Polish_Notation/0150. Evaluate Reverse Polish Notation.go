package main

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/":
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default:
			temp, _ := strconv.Atoi(tokens[i])
			stack = append(stack, temp)
		}
	}
	return stack[0]
}
