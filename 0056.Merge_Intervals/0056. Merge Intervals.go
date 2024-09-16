package main

import "sort"

func merge(intervals [][]int) (ans [][]int) {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] }) // 对intervals依据区间起始位置大小进行升序排序
	left, right := intervals[0][0], intervals[0][1]
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] <= right { // 当前区间的起始值没有超过之前的最大值，即需要与之前的区间合并
			right = max(intervals[i][1], right) // 新的最大值为两者中较大者
		} else { // 当前区间最小值超过之前区间的最大值，无法合并
			ans = append(ans, []int{left, right})          // 先前区间加入结果
			left, right = intervals[i][0], intervals[i][1] // 重设区间最值
		}
	}
	ans = append(ans, []int{left, right}) // 添加最后一个区间
	return
}
