package main

import (
	"sort"
	"strconv"
)

func timeToInt(time string) int {
	h, _ := strconv.Atoi(time[:2])
	m, _ := strconv.Atoi(time[3:])
	return h*60 + m
}

func findMinDifference(timePoints []string) int {
	n := len(timePoints)
	times := make([]int, 0, n)
	for _, point := range timePoints {
		times = append(times, timeToInt(point))
	}
	sort.Ints(times)
	ans := times[0] + 1440 - times[n-1]
	for i := 1; i < n; i++ {
		ans = min(ans, times[i]-times[i-1])
	}
	return ans
}
