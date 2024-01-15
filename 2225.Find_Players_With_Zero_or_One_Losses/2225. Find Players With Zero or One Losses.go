package main

import "sort"

func findWinners(matches [][]int) [][]int {
	var ans, winCnt, loseCnt = make([][]int, 2), make(map[int]int), make(map[int]int)
	for _, match := range matches {
		winCnt[match[0]]++
		loseCnt[match[1]]++
	}
	for k := range winCnt {
		if loseCnt[k] == 0 {
			ans[0] = append(ans[0], k)
		}
	}
	for k, v := range loseCnt {
		if v == 1 {
			ans[1] = append(ans[1], k)
		}
	}
	sort.Ints(ans[0])
	sort.Ints(ans[1])
	return ans
}
