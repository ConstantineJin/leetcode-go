package main

import "sort"

func maximumBags(capacity, rocks []int, additionalRocks int) (ans int) {
	for i, rock := range rocks {
		capacity[i] -= rock
	}
	sort.Ints(capacity)
	for i, c := range capacity {
		if c <= additionalRocks {
			additionalRocks -= c
		} else {
			return i
		}
	}
	return len(capacity)
}
