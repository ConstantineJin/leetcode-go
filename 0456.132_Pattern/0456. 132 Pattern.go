package main

import "math"

func find132pattern(nums []int) bool {
	n := len(nums)
	mx := math.MinInt             // "2"的最大值，初始化为 -∞
	st := []int{nums[n-1]}        // 使用单调栈维护可以作为“3”的元素，越靠近栈顶元素越小
	for i := n - 2; i >= 0; i-- { // 倒序枚举“1”
		if nums[i] < mx { // 如果当前元素小于“2”的最大值，那么就找到了“132”模式
			return true
		}
		for ; len(st) > 0 && nums[i] > st[len(st)-1]; st = st[:len(st)-1] { // 不断弹出栈顶元素直到当前元素不大于栈顶元素
			mx = st[len(st)-1] // 单调栈中元素弹出后可以成为“2”
		}
		if nums[i] > mx { // 当前元素可以作为“3”
			st = append(st, nums[i])
		}
	}
	return false
}
