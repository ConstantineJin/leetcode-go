package main

import (
	"slices"
	"strings"
)

func reversePrefix(word string, ch byte) string {
	var idx = strings.IndexByte(word, ch)
	if idx == -1 {
		return word
	}
	var res = []byte(word[:idx+1])
	slices.Reverse(res)
	return string(res) + word[idx+1:]
}
