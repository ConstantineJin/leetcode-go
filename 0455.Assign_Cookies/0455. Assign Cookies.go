package main

import "sort"

func findContentChildren(g []int, s []int) (ans int) {
	sort.Ints(g)
	sort.Ints(s)
	for i, j := 0, 0; i < len(g) && j < len(s); i++ {
		for j < len(s) && g[i] > s[j] {
			j++
		}
		if j < len(s) {
			ans++
			j++
		}
	}
	return
}
