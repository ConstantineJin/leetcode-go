package main

var mp = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) (ans int) {
	var n = len(s)
	for i := range s {
		var value = mp[s[i]]
		if i < n-1 && value < mp[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return
}
