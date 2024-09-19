package main

func longestContinuousSubstring(s string) int {
	ans, cnt := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i]-s[i-1] == 1 {
			cnt++
			ans = max(ans, cnt)
		} else {
			cnt = 1
		}
	}
	return ans
}
