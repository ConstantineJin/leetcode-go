package main

import "strings"

func areSentencesSimilar(sentence1, sentence2 string) bool {
	words1, words2 := strings.Split(sentence1, " "), strings.Split(sentence2, " ")
	i, j, m, n := 0, 0, len(words1), len(words2)
	for i < m && i < n && words1[i] == words2[i] {
		i++
	}
	for j < m-i && j < n-i && words1[m-j-1] == words2[n-j-1] {
		j++
	}
	return i+j == min(m, n)
}
