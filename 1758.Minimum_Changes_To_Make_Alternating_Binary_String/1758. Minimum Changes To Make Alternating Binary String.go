package main

func minOperations(s string) (ans int) {
	var str = []byte(s)
	var cnt int
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			if str[i] == '1' {
				str[i] = '0'
			} else {
				str[i] = '1'
			}
			cnt++
		}
	}
	ans = cnt
	str, cnt = []byte(s), 1
	if str[0] == '1' {
		str[0] = '0'
	} else {
		str[0] = '1'
	}
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			if str[i] == '1' {
				str[i] = '0'
			} else {
				str[i] = '1'
			}
			cnt++
		}
	}
	return min(ans, cnt)
}
