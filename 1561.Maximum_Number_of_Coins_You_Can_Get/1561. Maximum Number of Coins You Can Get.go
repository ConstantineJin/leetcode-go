package main

import "slices"

func maxCoins(piles []int) (ans int) {
	slices.Sort(piles)
	n := len(piles) / 3
	for i := 0; i < n; i++ {
		ans += piles[len(piles)-i*2-2]
	}
	return
}
