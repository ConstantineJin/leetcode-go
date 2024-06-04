package main

func longestPalindrome(s string) (ans int) {
	var mp = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mp[s[i]]++
	}
	var hasOdd int
	for _, v := range mp {
		if v%2 == 1 {
			hasOdd = 1
			ans += v - 1
		} else {
			ans += v
		}
	}
	ans += hasOdd
	return
}
