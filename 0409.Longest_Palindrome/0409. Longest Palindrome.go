package main

func longestPalindrome(s string) (ans int) {
	var mp = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mp[s[i]]++
	}
	var odd bool
	for _, v := range mp {
		if v%2 == 1 {
			odd = true
			ans += v - 1
		} else {
			ans += v
		}
	}
	if odd {
		ans += 1
	}
	return
}
