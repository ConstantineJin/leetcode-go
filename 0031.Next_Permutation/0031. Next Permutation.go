package main

import "sort"

func nextPermutation(nums []int) {
	var n = len(nums)
	var res, i = make([]int, n), n - 2
	for ; i >= 0; i-- { // 倒序遍历，寻找出现的第一个比右侧元素小的元素下标i
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i == -1 { // 原数组是降序排列
		sort.Ints(nums) // 改为升序排列
		return
	}
	var j = n - 1
	for ; j >= 0; j-- { // 再次倒序遍历，出现的第一个比nums[i]大的元素
		if nums[j] > nums[i] {
			break
		}
	}
	copy(res[:i], nums[:i])             // 前半段不用动
	nums[i], nums[j] = nums[j], nums[i] // 交换下标为i和j的元素位置
	sort.Ints(nums[i+1:])               // 后半段升序排列
	copy(res[i:], nums[i:])
}
