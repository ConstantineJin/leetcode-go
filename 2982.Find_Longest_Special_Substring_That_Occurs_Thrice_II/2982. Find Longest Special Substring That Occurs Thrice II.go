package main

func maximumLength(s string) (ans int) {
	var mem [26][3]int
	cur := 1
	for i, c := range s {
		if i == len(s)-1 || s[i] != s[i+1] {
			idx := c - 'a'
			if cur > mem[idx][0] {
				mem[idx][2] = mem[idx][1]
				mem[idx][1] = mem[idx][0]
				mem[idx][0] = cur
			} else if cur > mem[idx][1] {
				mem[idx][2] = mem[idx][1]
				mem[idx][1] = cur
			} else if cur > mem[idx][2] {
				mem[idx][2] = cur
			}
			cur = 1
		} else {
			cur++
		}
	}
	for _, top := range mem {
		ans = max(ans, top[0]-2, min(top[0]-1, top[1]), top[2])
	}
	if ans == 0 {
		return -1
	}
	return
}
