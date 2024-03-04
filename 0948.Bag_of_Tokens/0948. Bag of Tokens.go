package main

import "sort"

func bagOfTokensScore(tokens []int, power int) (ans int) {
	var score int
	sort.Ints(tokens)
	for len(tokens) > 0 {
		if power >= tokens[0] {
			power -= tokens[0]
			score++
			ans, tokens = max(ans, score), tokens[1:]
		} else if score > 0 {
			score--
			power += tokens[len(tokens)-1]
			tokens = tokens[:len(tokens)-1]
		} else {
			break
		}
	}
	return
}
