package main

import "slices"

func reverseStr(s string, k int) string {
	ss := []byte(s)
	n := len(s)
	for i := 0; i < n; i += k * 2 {
		slices.Reverse(ss[i:min(i+k, n)])
	}
	return string(ss)
}
