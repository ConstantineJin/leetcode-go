package main

func maxPower(s string) (ans int) {
	n, i, j := len(s), 0, 0
	for i < n {
		for j = i + 1; j < n && s[j] == s[i]; j++ {
		}
		ans, i = max(ans, j-i), j
	}
	return
}
