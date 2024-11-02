package main

import "strings"

func isCircularSentence(sentence string) bool {
	words := strings.Split(sentence, " ")
	n := len(words)
	words = append(words, words[0])
	for i, word := range words[:n] {
		if word[len(word)-1] != words[i+1][0] {
			return false
		}
	}
	return true
}
