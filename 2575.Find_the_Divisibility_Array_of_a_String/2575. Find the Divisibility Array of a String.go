package main

func divisibilityArray(word string, m int) []int {
	var n = len(word)
	var ans = make([]int, n)
	var temp int
	for i, c := range word {
		temp = (temp*10 + int(c) - '0') % m
		if temp == 0 {
			ans[i] = 1
		}
	}
	return ans
}
