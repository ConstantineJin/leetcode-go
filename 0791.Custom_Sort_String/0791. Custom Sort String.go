package main

import "strings"

func customSortString(order string, s string) (ans string) {
	var cnt [26]int
	for _, c := range s {
		cnt[c-'a']++
	}
	for _, c := range order {
		ans = ans + strings.Repeat(string(c), cnt[c-'a'])
		cnt[c-'a'] = 0
	}
	for i, v := range cnt {
		ans += strings.Repeat(string(rune('a'+i)), v)
	}
	return
}
