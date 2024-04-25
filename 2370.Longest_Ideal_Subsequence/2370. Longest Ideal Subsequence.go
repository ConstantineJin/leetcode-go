package main

import "slices"

func longestIdealString(s string, k int) int {
	var f [26]int
	for _, ch := range s {
		c := int(ch - 'a')
		f[c] = 1 + slices.Max(f[max(c-k, 0):min(c+k+1, 26)])
	}
	return slices.Max(f[:])
}
