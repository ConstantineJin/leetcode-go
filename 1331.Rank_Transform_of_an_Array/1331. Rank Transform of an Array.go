package main

import "slices"

type pair struct{ index, value int }

func arrayRankTransform(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return nil
	}
	pairs := make([]pair, 0, n)
	for i, v := range arr {
		pairs = append(pairs, pair{i, v})
	}
	slices.SortFunc(pairs, func(a, b pair) int { return a.value - b.value })
	ans := make([]int, n)
	rank := 1
	for i, p := range pairs {
		if i > 0 && p.value > pairs[i-1].value {
			rank++
		}
		ans[pairs[i].index] = rank
	}
	return ans
}
