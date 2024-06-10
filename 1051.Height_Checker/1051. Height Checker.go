package main

import "slices"

func heightChecker(heights []int) (ans int) {
	expected := slices.Clone(heights)
	slices.Sort(expected)
	for i, height := range heights {
		if height != expected[i] {
			ans++
		}
	}
	return
}
