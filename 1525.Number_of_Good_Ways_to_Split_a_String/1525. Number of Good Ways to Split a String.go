package main

func numSplits(s string) (ans int) {
	var left, right = make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s); i++ {
		right[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		left[s[i]]++
		right[s[i]]--
		if right[s[i]] == 0 {
			delete(right, s[i])
		}
		if len(left) == len(right) {
			ans++
		}
	}
	return
}
