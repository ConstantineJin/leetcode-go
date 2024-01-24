package main

func maximumNumberOfStringPairs(words []string) (ans int) {
	var seen = [26][26]bool{}
	for _, s := range words {
		var x, y = s[0] - 'a', s[1] - 'a'
		if seen[y][x] {
			ans++ // s 和 seen 中的 y+x 匹配
		} else {
			seen[x][y] = true
		}
	}
	return
}
