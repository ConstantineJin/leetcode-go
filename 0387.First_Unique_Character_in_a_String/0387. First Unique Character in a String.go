package main

func firstUniqChar(s string) int {
	var n = len(s)
	var pos [26][]int
	for i := 0; i < n; i++ {
		pos[s[i]-'a'] = append(pos[s[i]-'a'], i)
	}
	var ans = n
	for _, p := range pos {
		if len(p) == 1 {
			ans = min(ans, p[0])
		}
	}
	if ans == n {
		return -1
	} else {
		return ans
	}
}
