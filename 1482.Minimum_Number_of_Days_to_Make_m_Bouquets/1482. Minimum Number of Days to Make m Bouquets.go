package main

import (
	"slices"
	"sort"
)

func minDays(bloomDay []int, m int, k int) int {
	n := len(bloomDay)
	if n < m*k {
		return -1
	}
	minDay, maxDay := slices.Min(bloomDay), slices.Max(bloomDay)
	return minDay + sort.Search(maxDay-minDay, func(day int) bool { // 二分查找
		day += minDay
		var flowers, bouquets int
		for _, d := range bloomDay { // 遍历所有花
			if d > day { // 这束花这天还没开放
				flowers = 0
			} else {
				flowers++
				if flowers == k {
					bouquets++
					flowers = 0
				}
			}
		}
		return bouquets >= m
	})
}
