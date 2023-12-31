package main

import "math"

func maxLengthBetweenEqualCharacters(s string) int {
	var pos = [26][2]int{} // 每个字母首次和末次（不含首次）出现的位置
	for i := range pos {
		pos[i][0], pos[i][1] = -1, -1 // 初始化为-1，表示该字母未出现过
	}
	for i := 0; i < len(s); i++ {
		if pos[s[i]-'a'][0] == -1 { // 该字母是首次出现
			pos[s[i]-'a'][0] = i
		} else { // 该字母不是首次出现
			pos[s[i]-'a'][1] = i
		}
	}
	var ans = math.MinInt
	for _, p := range pos {
		if p[1] != -1 { // 当前字母出现至少两次
			ans = max(ans, p[1]-p[0])
		}
	}
	if ans == math.MinInt { // 没有任何字母出现两次
		return -1
	}
	return ans - 1 // 两个相同字符之间的子字符串长度是下标之差-1
}
