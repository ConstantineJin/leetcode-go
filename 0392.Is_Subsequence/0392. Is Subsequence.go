package main

func isSubsequence(s string, t string) bool {
	var m, n = len(s), len(t)
	if m > n {
		return false
	}
	var i int
	for j := 0; i < m && j < n; j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == m
}
