package main

import "math"

type pair struct {
	word string
	cost int
}

func minimumCost(target string, words []string, costs []int) int {
	var start [26][]pair
	for i, word := range words {
		start[word[0]-'a'] = append(start[word[0]-'a'], pair{word, costs[i]})
	}
	n := len(target)
	mem := make([]int, n)
	var dfs func(i int) int
	dfs = func(i int) int {
		if i == n {
			return 0
		}
		if mem[i] != 0 {
			return mem[i]
		}
		if len(start[target[i]-'a']) == 0 {
			mem[i] = -1
			return -1
		}
		res := math.MaxInt
		defer func() { mem[i] = res }()
		for _, p := range start[target[i]-'a'] {
			word, cost := p.word, p.cost
			if i+len(word) <= n && word == target[i:i+len(word)] {
				v := dfs(i + len(word))
				if v == -1 {
					continue
				}
				res = min(res, v+cost)
			}
		}
		if res == math.MaxInt {
			res = -1
		}
		return res
	}
	return dfs(0)
}
