package main

import "strings"

func removeKdigits(num string, k int) (ans string) {
	var stack []byte
	for i := range num {
		for k > 0 && len(stack) > 0 && num[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	ans = strings.TrimLeft(string(stack[:len(stack)-k]), "0")
	if ans == "" {
		return "0"
	}
	return
}
