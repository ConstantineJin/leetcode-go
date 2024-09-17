package main

import (
	"slices"
	"strings"
)

func uncommonFromSentences(s1, s2 string) (ans []string) {
	mp := make(map[string]int)
	for _, word := range slices.Concat(strings.Split(s1, " "), strings.Split(s2, " ")) {
		mp[word]++
	}
	for word, cnt := range mp {
		if cnt == 1 {
			ans = append(ans, word)
		}
	}
	return
}
