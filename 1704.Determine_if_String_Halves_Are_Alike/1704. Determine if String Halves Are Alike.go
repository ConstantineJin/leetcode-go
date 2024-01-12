package main

var vowel = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}

func halvesAreAlike(s string) bool {
	var cnt int
	var n = len(s)
	for i := 0; i < n/2; i++ {
		if vowel[s[i]] {
			cnt++
		}
	}
	for i := n / 2; i < n; i++ {
		if vowel[s[i]] {
			cnt--
		}
	}
	return cnt == 0
}
