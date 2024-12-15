package main

import "slices"

func minSetSize(arr []int) (ans int) {
	mp := make(map[int]int)
	for _, v := range arr {
		mp[v]++
	}
	freq := make([]int, 0, len(mp))
	for _, v := range mp {
		freq = append(freq, v)
	}
	slices.SortFunc(freq, func(a, b int) int { return b - a })
	var cnt, i int
	for cnt < len(arr)/2 {
		cnt += freq[i]
		ans++
		i++
	}
	return
}
