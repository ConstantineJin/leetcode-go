package main

func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	var n, i = len(intervals), 0
	for ; i < n && intervals[i][1] < newInterval[0]; i++ { // 第一阶段：原区间在新区间左侧且无重叠
		ans = append(ans, intervals[i]) // 无需合并
	}
	for ; i < n && intervals[i][0] <= newInterval[1]; i++ { // 第二阶段：原区间与新区间有重叠
		newInterval[0], newInterval[1] = min(newInterval[0], intervals[i][0]), max(newInterval[1], intervals[i][1]) // 不断合并原区间与新区间
	}
	ans = append(ans, newInterval) // 将合并后的新区间保存
	for ; i < n; i++ {             // 第三阶段：原区间在新区间右侧且无重叠
		ans = append(ans, intervals[i]) // 无需合并
	}
	return
}
