package main

import "strconv"

func longestCommonPrefix(arr1, arr2 []int) (ans int) {
	has := make(map[string]bool)
	for _, num := range arr1 {
		s := strconv.Itoa(num)
		for i := 1; i <= len(s); i++ {
			has[s[:i]] = true
		}
	}
	for _, num := range arr2 {
		s := strconv.Itoa(num)
		for i := 1; i <= len(s) && has[s[:i]]; i++ {
			ans = max(ans, i)
		}
	}
	return
}
