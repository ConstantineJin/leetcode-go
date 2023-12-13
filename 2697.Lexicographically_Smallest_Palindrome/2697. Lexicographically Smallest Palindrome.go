package main

func makeSmallestPalindrome(s string) string {
	var n = len(s)
	var ans = make([]byte, n)
	for i := 0; i <= n/2; i++ {
		if s[i] <= s[n-i-1] {
			ans[i], ans[n-i-1] = s[i], s[i]
		} else {
			ans[i], ans[n-i-1] = s[n-i-1], s[n-i-1]
		}
	}
	return string(ans)
}
