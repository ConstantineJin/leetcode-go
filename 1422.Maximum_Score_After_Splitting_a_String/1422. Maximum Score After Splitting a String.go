package main

func maxScore(s string) (ans int) {
	var n = len(s)
	var l, r int
	for _, v := range s {
		if v == '1' {
			r++
		}
	}
	for i, v := range s {
		if v == '0' {
			l++
		} else {
			r--
		}
		if i != n-1 {
			ans = max(ans, l+r)
		}
	}
	return
}
