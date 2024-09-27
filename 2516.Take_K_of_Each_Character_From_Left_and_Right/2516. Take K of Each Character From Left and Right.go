package main

func takeCharacters(s string, k int) int {
	var count [3]int
	for _, c := range s {
		count[c-'a']++
	}
	for _, cnt := range count {
		if cnt < k {
			return -1
		}
	}
	var mx, left int
	for right, c := range s {
		c -= 'a'
		count[c]--
		for count[c] < k {
			count[s[left]-'a']++
			left++
		}
		mx = max(mx, right-left+1)
	}
	return len(s) - mx
}
