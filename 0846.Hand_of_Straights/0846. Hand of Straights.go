package main

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	cnt := make(map[int]int)
	for _, v := range hand {
		cnt[v]++
	}
	sort.Ints(hand)
	for _, x := range hand {
		if cnt[x] == 0 {
			continue
		}
		for i := x; i < x+groupSize; i++ {
			if cnt[i] == 0 {
				return false
			}
			cnt[i]--
		}
	}
	return true
}
