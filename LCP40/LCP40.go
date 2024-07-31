package main

import "slices"

func maxmiumScore(cards []int, cnt int) (ans int) {
	slices.SortFunc(cards, func(a, b int) int { return b - a })
	var s int
	odd, even := -1, -1
	for _, card := range cards[:cnt] {
		s += card
		if card%2 == 0 {
			even = card
		} else {
			odd = card
		}
	}
	if s%2 == 0 {
		return s
	}
	for _, card := range cards[cnt:] {
		if card%2 == 1 && even != -1 {
			ans = max(ans, s-even+card)
		} else if card%2 == 0 && odd != -1 {
			ans = max(ans, s-odd+card)
		}
	}
	return
}
