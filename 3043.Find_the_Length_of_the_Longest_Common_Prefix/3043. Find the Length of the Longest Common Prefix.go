package main

import (
	"strconv"
)

func longestCommonPrefix(arr1, arr2 []int) (ans int) {
	var has = make(map[string]bool)
	for _, x := range arr1 {
		var s = strconv.Itoa(x)
		for i := 1; i <= len(s); i++ {
			has[s[:i]] = true
		}
	}

	for _, x := range arr2 {
		var s = strconv.Itoa(x)
		for i := 1; i <= len(s) && has[s[:i]]; i++ {
			ans = max(ans, i)
		}
	}
	return
}
