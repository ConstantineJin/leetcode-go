package main

import "math"

func myAtoi(s string) (ans int) {
	pos, i := true, 0
	for i = 0; i < len(s) && s[i] == ' '; i++ {
	}
	if i == len(s) {
		return
	}
	switch s[i] {
	case '+':
		i++
	case '-':
		pos = false
		i++
	}
	for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
		if !pos && (ans > math.MaxInt32/10 || ans == math.MaxInt32/10 && int(s[i])-'0' > 8) {
			return math.MinInt32
		} else if pos && (ans > math.MaxInt32/10 || ans == math.MaxInt32/10 && int(s[i])-'0' > 7) {
			return math.MaxInt32
		}
		ans = ans*10 + int(s[i]) - '0'
	}
	if pos {
		return
	}
	return -ans
}
