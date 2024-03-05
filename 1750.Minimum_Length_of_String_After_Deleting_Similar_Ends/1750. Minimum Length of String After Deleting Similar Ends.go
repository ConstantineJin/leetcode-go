package main

func minimumLength(s string) int {
	var i, j = 0, len(s) - 1
	for i < j && s[i] == s[j] {
		var c = s[i]
		for ; i <= j && s[i] == c; i++ {
		}
		for ; i <= j && s[j] == c; j-- {
		}
	}
	return j - i + 1
}
