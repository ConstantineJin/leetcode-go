package main

func countSubstrings(s string) (ans int) {
	var n = len(s)
	for i := 0; i < n; i++ {
		for j := 0; j <= i && i+j < n; j++ {
			if s[i-j] == s[i+j] {
				ans++
			} else {
				break
			}
		}
	}
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			for j := 0; j <= i-1 && i+j < n; j++ {
				if s[i-j-1] == s[i+j] {
					ans++
				} else {
					break
				}
			}
		}
	}
	return
}
