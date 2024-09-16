package main

var mp = map[rune]int{'a': 0, 'e': 1, 'i': 2, 'o': 3, 'u': 4}

func findTheLongestSubstring(s string) (ans int) {
	var mask int             // 使用 mask 维护每个原因出现次数的奇偶性
	pos := make([]int, 1<<5) // pos[mask] 维护五个元音第一次出现奇偶性为 mask 时的下标
	for i := range pos {
		pos[i] = -1
	}
	pos[0] = 0
	for i, c := range s {
		if v, ok := mp[c]; ok {
			mask ^= 1 << v
		}
		if pos[mask] >= 0 {
			ans = max(ans, i+1-pos[mask])
		} else {
			pos[mask] = i + 1
		}
	}
	return ans
}
