package main

func getSmallestString(s string) string {
	ans := []byte(s)
	for i := 1; i < len(s); i++ {
		if s[i]%2 == s[i-1]%2 && s[i-1] > s[i] {
			ans[i-1], ans[i] = ans[i], ans[i-1]
			break
		}
	}
	return string(ans)
}
