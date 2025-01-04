package main

func countPalindromicSubsequence(s string) (ans int) {
	var vis [26][26]bool
	var first, last [26]int
	n := len(s)
	for i := range 26 {
		first[i], last[i] = n, -1
	}
	for i, c := range s {
		first[c-'a'], last[c-'a'] = min(first[c-'a'], i), max(last[c-'a'], i)
	}
	for i, left := range first {
		if left == n {
			continue
		}
		for j := left + 1; j < last[i]; j++ {
			if !vis[i][s[j]-'a'] {
				vis[i][s[j]-'a'] = true
				ans++
			}
		}
	}
	return
}
