package main

import "slices"

func numberOfWeeks(milestones []int) int64 {
	var sum int
	for _, milestone := range milestones {
		sum += milestone
	}
	mx := slices.Max(milestones)
	if mx > sum-mx+1 { // 如果最大的元素大到比其余所有元素只和再+1还大，那么最大的元素无法完成
		return int64((sum-mx)*2 + 1)
	}
	return int64(sum) // 否则一定可以完成所有任务
}
