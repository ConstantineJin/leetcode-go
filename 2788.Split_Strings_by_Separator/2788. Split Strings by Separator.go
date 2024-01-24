package main

import "strings"

func splitWordsBySeparator(words []string, separator byte) (ans []string) {
	for _, word := range words {
		for _, w := range strings.Split(word, string(separator)) {
			if len(w) != 0 {
				ans = append(ans, w)
			}
		}
	}
	return
}
