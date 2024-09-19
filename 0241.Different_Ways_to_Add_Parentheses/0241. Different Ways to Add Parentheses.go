package main

import "strconv"

func diffWaysToCompute(expression string) (ans []int) {
	num, err := strconv.Atoi(expression)
	if err == nil { // 如果当前表达式只含有数字，直接返回
		return []int{num}
	}
	for i, c := range expression {
		if c == '+' || c == '-' || c == '*' { // 如果当前字符为运算符，则递归计算左半边和右半边的计算结果
			left, right := diffWaysToCompute(expression[:i]), diffWaysToCompute(expression[i+1:])
			for _, leftNum := range left {
				for _, rightNum := range right {
					switch c {
					case '+':
						ans = append(ans, leftNum+rightNum)
					case '-':
						ans = append(ans, leftNum-rightNum)
					case '*':
						ans = append(ans, leftNum*rightNum)
					}
				}
			}
		}
	}
	return
}
