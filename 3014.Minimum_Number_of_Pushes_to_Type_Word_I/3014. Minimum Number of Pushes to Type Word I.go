package main

import "sort"

func minimumPushes(word string) (ans int) {
	var cnt = make([]int, 26)
	for _, ch := range word {
		cnt[ch-'a']++
	}
	sort.Ints(cnt)
	for i := 25; i >= 0; i-- {
		ans += (33 - i) / 8 * cnt[i]
	}
	return
}
