package main

func shiftingLetters(s string, shifts [][]int) string {
	diffs := make([]int, len(s)+1) // 差分数组
	for _, shift := range shifts {
		start, end, move := shift[0], shift[1], shift[2]*2-1
		diffs[start] += move
		diffs[end+1] -= move
	}
	ans, shift := []byte(s), 0
	for i, c := range ans {
		shift = (shift+diffs[i])%26 + 26
		ans[i] = (c-'a'+byte(shift))%26 + 'a'
	}
	return string(ans)
}
