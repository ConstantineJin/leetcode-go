package main

import (
	"slices"
	"sort"
)

func minimizedMaximum(n int, quantities []int) int {
	return sort.Search(slices.Max(quantities), func(i int) bool {
		var res int
		for _, quantity := range quantities {
			res += (quantity + i) / (i + 1)
		}
		return res <= n
	}) + 1
}
