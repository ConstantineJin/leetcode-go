package main

func minAddToMakeValid(s string) (ans int) {
	var cnt int
	for _, c := range s {
		if c == '(' {
			cnt++
		} else if cnt > 0 {
			cnt--
		} else {
			ans++
		}
	}
	return ans + cnt
}
