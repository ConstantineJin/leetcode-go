package main

import "math"

func longestCommonPrefix(strs []string) (ans string) {
	var l = math.MaxInt
	for _, str := range strs {
		l = min(l, len(str))
	}
	for i := 0; i < l; i++ {
		var flag = true
		for _, str := range strs {
			if str[i] != strs[0][i] {
				flag = false
				break
			}
		}
		if flag {
			ans += string(strs[0][i])
		} else {
			break
		}
	}
	return
}
