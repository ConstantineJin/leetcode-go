package main

func canSeePersonsCount(heights []int) []int {
	var n = len(heights)
	var ans = make([]int, n)
	var stack = []int{n - 1}      // 单调栈
	ans[n-1] = 0                  // 最后一个人的右边没有人
	for i := n - 2; i >= 0; i-- { // 倒序遍历heights
		for j := len(stack) - 1; j >= 0 && heights[i] > heights[stack[len(stack)-1]]; j-- { // 栈不为空且当前元素大于栈顶元素时
			stack = stack[:len(stack)-1] // 栈顶元素出栈
			ans[i]++                     // 栈顶元素一定能被i看到
		}
		if len(stack) != 0 { // 如果此时栈不为空
			ans[i]++ // 说明i右侧存在至少一个更高的元素，那么i还能再看到一个人
		}
		stack = append(stack, i) // 当前元素入栈
	}
	return ans
}
