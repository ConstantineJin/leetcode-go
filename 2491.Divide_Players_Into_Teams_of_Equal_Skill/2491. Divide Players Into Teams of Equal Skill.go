package main

import "sort"

func dividePlayers(skill []int) (ans int64) {
	sort.Ints(skill)
	n := len(skill)
	for i, s := 0, skill[0]+skill[n-1]; i < n/2; i++ {
		if skill[i]+skill[n-i-1] != s {
			return -1
		}
		ans += int64(skill[i] * skill[n-i-1])
	}
	return
}
