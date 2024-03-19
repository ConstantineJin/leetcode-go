package main

func largestRectangleArea(heights []int) (ans int) {
	var n = len(heights)
	var left, right = make([]int, n), make([]int, n) // left[i]/right[i]表示heights[i]的左/右侧第一个比heights低的下标，如果不存在则为-1/n
	var st []int
	for i, height := range heights {
		for len(st) > 0 && height <= heights[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			left[i] = -1
		} else {
			left[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	st = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && heights[i] <= heights[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			right[i] = n
		} else {
			right[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	for i, height := range heights {
		ans = max(ans, height*(right[i]-left[i]-1))
	}
	return ans
}
