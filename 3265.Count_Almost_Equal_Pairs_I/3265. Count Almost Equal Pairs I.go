package main

import (
	"strconv"
	"strings"
)

func isSimilar(x, y int) bool {
	if x == y {
		return true
	}
	if x > y {
		x, y = y, x
	}
	a, b := strconv.Itoa(x), strconv.Itoa(y)
	if len(a) < len(b) {
		a = strings.Repeat("0", len(b)-len(a)) + a
	}
	var diff int
	var c0, c1 byte
	for i := range a {
		if a[i] != b[i] {
			switch diff {
			case 0:
				c0, c1 = a[i], b[i]
			case 1:
				if a[i] != c1 || b[i] != c0 {
					return false
				}
			case 2:
				return false
			}
			diff++
		}
	}
	return diff == 2
}

func countPairs(nums []int) (ans int) {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if isSimilar(num, nums[j]) {
				ans++
			}
		}
	}
	return
}
