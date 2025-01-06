package main

import "slices"

func maxConsecutive(bottom, top int, special []int) int {
	slices.Sort(special)
	ans := max(special[0]-bottom, top-special[len(special)-1])
	for i := 1; i < len(special); i++ {
		ans = max(ans, special[i]-special[i-1]-1)
	}
	return ans
}
