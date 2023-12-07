package main

func sumSubarrayMins(arr []int) (ans int) {
	arr = append(arr, -1)   // 末尾添加哨兵
	stack := []int{-1}      // 单调栈储存下标，添加哨兵
	for i, v := range arr { // 一次遍历
		for len(stack) > 1 && arr[stack[len(stack)-1]] >= v { // 栈不为空且栈顶元素大于等于当前遍历到的元素
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]                        // 栈顶元素出栈
			ans += arr[j] * (j - stack[len(stack)-1]) * (i - j) // 加权贡献，左边界j-stack[len(stack)-1]，右边界i-j
		}
		stack = append(stack, i) // 将当前元素下标入栈
	}
	return ans % (1e9 + 7)
}
