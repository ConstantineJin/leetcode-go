package main

import (
	"slices"
	"strings"
)

func stringMatching(words []string) (ans []string) {
	slices.SortFunc(words, func(a, b string) int { return len(a) - len(b) })
	for i, word0 := range words {
		for _, word1 := range words[i+1:] {
			if strings.Contains(word1, word0) {
				ans = append(ans, word0)
				break
			}
		}
	}
	return ans
}
