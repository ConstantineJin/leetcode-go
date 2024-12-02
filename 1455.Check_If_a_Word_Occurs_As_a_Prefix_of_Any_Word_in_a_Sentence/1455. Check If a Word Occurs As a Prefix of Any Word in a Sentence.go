package main

import "strings"

func isPrefixOfWord(sentence, searchWord string) int {
	words := strings.Split(sentence, " ")
	for i, word := range words {
		if strings.HasPrefix(word, searchWord) {
			return i + 1
		}
	}
	return -1
}
