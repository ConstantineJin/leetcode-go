package main

import "strings"

func countKConstraintSubstrings(s string, k int) (ans int) {
	n := len(s)
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if strings.Count(s[i:j], "0") <= k || strings.Count(s[i:j], "1") <= k {
				ans++
			}
		}
	}
	return
}
