package main

func minimumSteps(s string) int64 {
	var ans, zeroCnt int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			zeroCnt++
		} else {
			ans += zeroCnt
		}
	}
	return int64(ans)
}
