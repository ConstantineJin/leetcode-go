package main

import "sort"

func relocateMarbles(nums, moveFrom, moveTo []int) []int {
	pos := make(map[int]struct{})
	for _, num := range nums {
		pos[num] = struct{}{}
	}
	for i, from := range moveFrom {
		delete(pos, from)
		pos[moveTo[i]] = struct{}{}
	}
	ans := make([]int, 0, len(pos))
	for i := range pos {
		ans = append(ans, i)
	}
	sort.Ints(ans)
	return ans
}
