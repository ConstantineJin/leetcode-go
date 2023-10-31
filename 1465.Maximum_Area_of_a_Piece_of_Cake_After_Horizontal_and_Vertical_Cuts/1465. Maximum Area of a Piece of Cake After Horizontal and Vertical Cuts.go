package main

import "sort"

func findMax(arr []int, l int) int {
	sort.Ints(arr)
	m := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] > m {
			m = arr[i] - arr[i-1]
		}
	}
	if l-arr[len(arr)-1] > m {
		m = l - arr[len(arr)-1]
	}
	return m
}

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	horizontalMax, verticalMax := findMax(horizontalCuts, h), findMax(verticalCuts, w)
	return horizontalMax * verticalMax % (1e9 + 7)
}
