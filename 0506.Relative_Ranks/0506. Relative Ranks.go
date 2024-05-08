package main

import (
	"sort"
	"strconv"
)

type athlete struct{ index, score int }

func findRelativeRanks(score []int) []string {
	n := len(score)
	athletes, ans := make([]athlete, n), make([]string, n)
	for i, s := range score {
		athletes[i] = athlete{index: i, score: s}
	}
	sort.Slice(athletes, func(i, j int) bool { return athletes[i].score > athletes[j].score })
	rank := [3]string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	for i, a := range athletes {
		if i < 3 {
			ans[a.index] = rank[i]
		} else {
			ans[a.index] = strconv.Itoa(i + 1)
		}
	}
	return ans
}
