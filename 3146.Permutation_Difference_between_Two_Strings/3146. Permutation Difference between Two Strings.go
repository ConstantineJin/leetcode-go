package main

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findPermutationDifference(s, t string) (ans int) {
	mp := make(map[rune]int)
	for i, c := range s {
		mp[c] = i
	}
	for i, c := range t {
		ans += abs(i - mp[c])
	}
	return
}
