package main

func minGroups(intervals [][]int) (ans int) {
	var diff [1e6 + 1]int // 等价于求每个点最多被多少个区间覆盖。可以使用差分数组维护
	for _, interval := range intervals {
		diff[interval[0]-1]++
		diff[interval[1]]--
	}
	var sum int
	for _, d := range diff {
		sum += d
		ans = max(ans, sum)
	}
	return
}
