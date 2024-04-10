package main

import "sort"

func deckRevealedIncreasing(deck []int) []int {
	var n = len(deck)
	var ans = make([]int, n)
	sort.Ints(deck)
	var queue = make([]int, n)
	for i := 0; i < n; i++ {
		queue[i] = i
	}
	for i := 0; len(queue) > 0; i++ {
		ans[queue[0]] = deck[i]
		queue = queue[1:]
		if len(queue) > 0 {
			queue = append(queue, queue[0])[1:]
		}
	}
	return ans
}
