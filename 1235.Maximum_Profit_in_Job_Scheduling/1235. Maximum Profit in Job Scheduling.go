package main

import "sort"

type job struct{ start, end, profit int }

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	var n = len(profit)
	var jobs = make([]job, n)
	for i := range jobs {
		jobs[i] = job{start: startTime[i], end: endTime[i], profit: profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool { return jobs[i].end < jobs[j].end }) // 按照结束时间对工作排序
	var f = make([]int, n+1)                                                   // f[i]表示按照结束时间排序后的前i个工作的最大报酬
	for i, job := range jobs {
		var j = sort.Search(i, func(j int) bool { return jobs[j].end > job.start }) // j是最大的满足endTime[j]≤startTime[i]的j，不存在时为−1；为避免处理-1将所有f的下标都+1
		f[i+1] = max(f[i], f[j]+job.profit)                                         // 不选第i个工作：f[i]=f[i-1]；选第i个工作：f[i]=f[j]+profit[i]
	}
	return f[n]
}
