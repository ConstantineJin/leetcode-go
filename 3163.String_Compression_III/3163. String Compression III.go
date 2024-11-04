package main

import (
	"bytes"
	"strings"
)

func compressedString(word string) string {
	var builder strings.Builder
	builder.Grow(len(word) * 2)
	i := -1
	for j := range word {
		c := word[j]
		if j+1 == len(word) || c != word[j+1] {
			k := j - i
			builder.WriteString(string(bytes.Repeat([]byte{'9', c}, k/9)))
			if k%9 > 0 {
				builder.WriteString(string('0' + byte(k%9)))
				builder.WriteByte(c)
			}
			i = j
		}
	}
	return builder.String()
}
