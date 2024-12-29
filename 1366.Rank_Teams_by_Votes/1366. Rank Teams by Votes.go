package main

import (
	"cmp"
	"maps"
	"slices"
)

func rankTeams(votes []string) string {
	cnt := make(map[rune][]int)
	for _, c := range votes[0] {
		cnt[c] = make([]int, len(votes[0]))
	}
	for _, vote := range votes {
		for i, c := range vote {
			cnt[c][i]++
		}
	}
	return string(slices.SortedFunc(maps.Keys(cnt), func(a rune, b rune) int { return cmp.Or(slices.Compare(cnt[b], cnt[a]), cmp.Compare(a, b)) }))
}
