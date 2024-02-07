package main

import (
	"sort"
	"strings"
)

type pair struct {
	ch  uint8
	cnt int
}

func frequencySort(s string) string {
	var mp = make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		mp[s[i]]++
	}
	var pairs = make([]pair, len(mp))
	var i int
	for k, v := range mp {
		pairs[i] = pair{ch: k, cnt: v}
		i++
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].cnt > pairs[j].cnt
	})
	var ans []byte
	for _, v := range pairs {
		ans = append(ans, strings.Repeat(string(v.ch), v.cnt)...)
	}
	return string(ans)
}
