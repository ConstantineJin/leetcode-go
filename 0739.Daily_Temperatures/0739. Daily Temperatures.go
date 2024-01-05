package main

func dailyTemperatures(temperatures []int) []int {
	var n = len(temperatures)
	var ans = make([]int, n)
	var stack = []int{n - 1} // 单调栈
	ans[n-1] = 0             // 对于最后一天，显然不存在后续的更高温度
	for i := n - 2; i >= 0; i-- {
		for len(stack) != 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] { // 当栈不为空且当前温度不小于栈顶温度时
			stack = stack[:len(stack)-1] // 栈顶元素出栈（因为它不可能是前面任何一天的后续更高温度）
		}
		if len(stack) != 0 { // 如果栈不为空
			ans[i] = stack[len(stack)-1] - i // 栈顶就是下一个更高温度
		} // 如果栈为空，说明不存在后续更高温度
		stack = append(stack, i) // 当前日期入栈
	}
	return ans
}
