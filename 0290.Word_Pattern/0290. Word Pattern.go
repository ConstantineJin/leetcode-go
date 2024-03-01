package main

import "strings"

func wordPattern(pattern string, s string) bool {
	var p2s, s2p, words = make(map[byte]string), make(map[string]byte), strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	for i, word := range words {
		if _, ok := p2s[pattern[i]]; ok && p2s[pattern[i]] != word {
			return false
		}
		if _, ok := s2p[word]; ok && s2p[word] != pattern[i] {
			return false
		}
		p2s[pattern[i]], s2p[word] = word, pattern[i]
	}
	return true
}
