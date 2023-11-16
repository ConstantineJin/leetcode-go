package main

import (
	"strconv"
)

func compress(chars []byte) int {
	n, i, j, k := len(chars), 0, 0, 0
	for i < n {
		for j = i + 1; j < n && chars[j] == chars[i]; j++ {
		}
		chars[k] = chars[i]
		k++
		if j != i+1 {
			ch := []byte(strconv.Itoa(j - i))
			for l, c := range ch {
				chars[k+l] = c
			}
			k += len(ch)
		}
		i = j
	}
	return k
}
