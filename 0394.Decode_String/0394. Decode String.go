package main

import (
	"strconv"
	"strings"
)

func decodeString(s string) (res string) {
	var numStack []int
	var strStack []string
	var num int
	for _, char := range s {
		if char >= '0' && char <= '9' {
			var n, _ = strconv.Atoi(string(char))
			num = num*10 + n
		} else if char == '[' {
			strStack = append(strStack, res)
			res = ""
			numStack = append(numStack, num)
			num = 0
		} else if char == ']' { // 出栈
			res = strStack[len(strStack)-1] + strings.Repeat(res, numStack[len(numStack)-1])
			strStack, numStack = strStack[:len(strStack)-1], numStack[:len(numStack)-1]
		} else {
			res += string(char)
		}
	}
	return
}
