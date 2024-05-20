package main

func longestAwesome(s string) (ans int) {
	const D = 10
	n := len(s)
	pos := make([]int, 1<<D)
	for i := range pos {
		pos[i] = n
	}
	pos[0] = -1
	var pre int
	for i, c := range s {
		pre ^= 1 << (c - '0')
		for d := 0; d < D; d++ {
			ans = max(ans, i-pos[pre^(1<<d)])
		}
		ans = max(ans, i-pos[pre])
		if pos[pre] == n {
			pos[pre] = i
		}
	}
	return
}
