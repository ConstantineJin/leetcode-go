package main

func maximumSumOfHeights(maxHeights []int) int64 {
	var n = len(maxHeights)
	var suf = make([]int, n+1)
	var stack = []int{n} // 哨兵
	var sum, pre int
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 1 && maxHeights[i] <= maxHeights[stack[len(stack)-1]] {
			var j = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			sum -= maxHeights[j] * (stack[len(stack)-1] - j) // 撤销之前加到sum中的
		}
		sum += maxHeights[i] * (stack[len(stack)-1] - i) // 从i到stack[len(stack)-1]-1都是height
		suf[i], stack = sum, append(stack, i)
	}
	var ans = sum
	stack = []int{-1} // 哨兵
	for i, height := range maxHeights {
		for len(stack) > 1 && height <= maxHeights[stack[len(stack)-1]] {
			var j = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pre -= maxHeights[j] * (j - stack[len(stack)-1]) // 撤销之前加到pre中的
		}
		pre += height * (i - stack[len(stack)-1]) // 从stack[len(stack)-1]+1到i都是height
		ans, stack = max(ans, pre+suf[i+1]), append(stack, i)
	}
	return int64(ans)
}
