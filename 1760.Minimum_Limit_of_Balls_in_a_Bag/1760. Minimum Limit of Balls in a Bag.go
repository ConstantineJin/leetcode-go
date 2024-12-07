package main

import (
	"slices"
	"sort"
)

func minimumSize(nums []int, maxOperations int) int {
	return sort.Search(slices.Max(nums), func(i int) bool {
		if i == 0 {
			return false
		}
		var ops int
		for _, num := range nums {
			ops += (num - 1) / i
		}
		return ops <= maxOperations
	})
}
