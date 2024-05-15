package main

import "slices"

func findMinimumTime(tasks [][]int) (ans int) {
	slices.SortFunc(tasks, func(a, b []int) int { return a[1] - b[1] }) // 按照每个任务的结束时间升序排序
	run := make([]bool, tasks[len(tasks)-1][1]+1)                       // 记录每个时刻是否运行
	for _, task := range tasks {
		start, end, duration := task[0], task[1], task[2]
		for _, b := range run[start : end+1] { // 统计任务时间区间内已运行的时间点
			if b {
				duration--
			}
		}
		for i := end; duration > 0; i-- { // 尽量把新增的运行时间点安排在区间的后缀上
			if !run[i] {
				run[i] = true
				duration--
				ans++
			}
		}
	}
	return
}
