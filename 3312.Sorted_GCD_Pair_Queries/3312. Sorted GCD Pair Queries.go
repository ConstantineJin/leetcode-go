package main

import (
	"slices"
	"sort"
)

func gcdValues(nums []int, queries []int64) []int {
	mx := slices.Max(nums)
	freq := make([]int, mx+1)
	for _, num := range nums {
		freq[num]++
	}
	cntGcd := make([]int, mx+1)
	for i := mx; i > 0; i-- {
		var c int
		for j := i; j <= mx; j += i {
			c += freq[j]
			cntGcd[i] -= cntGcd[j]
		}
		cntGcd[i] += c * (c - 1) / 2 // 从 c 个数里选 2 个
	}
	for i := 2; i <= mx; i++ {
		cntGcd[i] += cntGcd[i-1] // 原地求前缀和
	}
	ans := make([]int, len(queries))
	for i, query := range queries {
		ans[i] = sort.SearchInts(cntGcd, int(query)+1)
	}
	return ans
}
