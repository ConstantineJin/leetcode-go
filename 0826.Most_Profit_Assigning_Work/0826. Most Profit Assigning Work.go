package main

import "slices"

type job struct{ difficulty, profit int }

func maxProfitAssignment(difficulty []int, profit []int, worker []int) (ans int) {
	n := len(difficulty)
	jobs := make([]job, n)
	for i := range difficulty {
		jobs[i] = job{difficulty: difficulty[i], profit: profit[i]}
	}
	slices.SortFunc(jobs, func(a, b job) int { return a.difficulty - b.difficulty })
	slices.Sort(worker)
	var mx, i int
	for _, w := range worker {
		for ; i < n && w >= jobs[i].difficulty; i++ {
			mx = max(mx, jobs[i].profit)
		}
		ans += mx
	}
	return
}
