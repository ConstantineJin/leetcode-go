package main

import "sort"

var arr []int

func init() {
next:
	for i := 1; i <= 1e4; i++ {
		x := i
		for x > 0 {
			d := x % 10
			if d == 0 || i%d != 0 {
				continue next
			}
			x /= 10
		}
		arr = append(arr, i)
	}
}

func selfDividingNumbers(left, right int) []int {
	return arr[sort.SearchInts(arr, left):sort.SearchInts(arr, right+1)]
}
