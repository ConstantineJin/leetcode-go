package main

// 单调栈
func nextGreaterElements(nums []int) []int {
	var n, m, idx = len(nums), nums[0], 0 // 最大元素的值和下标
	for i, num := range nums {
		if num > m { // 只保留最前面的最大元素
			m, idx = num, i
		}
	}
	var ans, st = make([]int, n), make([]int, 0)
	nums = append(nums, nums[:idx+1]...) // 模拟循环，只需要把到第一个最大元素（含）之前的nums子数组添加到nums之后
	for i := len(nums) - 1; i >= 0; i-- {
		for len(st) > 0 && nums[i] >= st[len(st)-1] { // 栈不为空且当前元素不小于栈顶元素时
			st = st[:len(st)-1] // 栈顶元素出栈
		}
		if i < n {
			if len(st) == 0 { // 栈为空
				ans[i] = -1 // 说明不存在更大元素
			} else { // 栈不空
				ans[i] = st[len(st)-1] // 下一个更大元素就是栈顶元素
			}
		}
		st = append(st, nums[i]) // 当前元素入栈
	}
	return ans
}
