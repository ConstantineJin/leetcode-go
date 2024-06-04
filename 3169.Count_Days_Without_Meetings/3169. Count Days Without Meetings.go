package main

import "sort"

func countDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][0] < meetings[j][0] })
	n := len(meetings)
	merged := make([][]int, 1)
	merged[0] = meetings[0]
	for i := 1; i < n; i++ {
		if meetings[i][0] <= merged[len(merged)-1][1] {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], meetings[i][1])
		} else {
			merged = append(merged, meetings[i])
		}
	}
	for _, meeting := range merged {
		days -= meeting[1] - meeting[0] + 1
	}
	return days
}
