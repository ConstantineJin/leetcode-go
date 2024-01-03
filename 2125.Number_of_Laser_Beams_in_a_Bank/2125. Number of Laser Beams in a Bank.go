package main

import (
	"strings"
)

func numberOfBeams(bank []string) (ans int) {
	var pre, cur int
	for _, s := range bank {
		cur = strings.Count(s, "1")
		if cur != 0 {
			ans += pre * cur
			pre = cur
		}
	}
	return
}
