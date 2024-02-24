package main

func lengthOfLastWord(s string) int {
	var i, j int
	for i = len(s) - 1; i >= 0 && s[i] == ' '; i-- {
	}
	for j = i; j >= 0 && s[j] != ' '; j-- {
	}
	return i - j
}
