package main

import (
	"slices"
	"sort"
)

// 有n门考试，第i门考试需要nums[i]复习；一天只能复习一门课程；在第j天可以参加第changeIndices[j]门课程的考试，且当天不能复习。至少需要多少天？
func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
	var n, m = len(nums), len(changeIndices)
	if n > m {
		return -1
	}
	var ddl = make([]int, n)
	var ans = n + sort.Search(m+1-n, func(mx int) bool {
		mx += n
		clear(ddl)
		for t, idx := range changeIndices[:mx] { // 每门考试最晚考试时间
			ddl[idx-1] = t + 1
		}
		if slices.Contains(ddl, 0) { // 如果有课程没有考试时间
			return false
		}
		var cnt int
		for i, idx := range changeIndices[:mx] {
			idx--                // 原题里下标从1开始
			if i == ddl[idx]-1 { // 这门课程的最晚一次考试
				if nums[idx] > cnt { // 所需复习时间超过剩余天数
					return false
				}
				cnt -= nums[idx] // 复习这门考试
			} else {
				cnt++
			}
		}
		return true
	})
	if ans > m {
		return -1
	}
	return ans
}
