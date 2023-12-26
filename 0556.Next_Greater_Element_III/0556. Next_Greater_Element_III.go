package main

import (
	"math"
	"sort"
)

func nextGreaterElement(n int) (ans int) {
	var nums []int
	for n > 0 {
		nums = append([]int{n % 10}, nums...)
		n /= 10
	}
	var m = len(nums)
	var res, i = make([]int, m), m - 2
	for ; i >= 0; i-- { // 倒序遍历，寻找出现的第一个比右侧元素小的元素下标i
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i == -1 { // 不存在这样的元素
		return -1
	}
	var j = m - 1
	for ; j >= 0; j-- { // 再次倒序遍历，出现的第一个比nums[i]大的元素
		if nums[j] > nums[i] {
			break
		}
	}
	copy(res[:i], nums[:i])             // 前半段不用动
	nums[i], nums[j] = nums[j], nums[i] // 交换下标为i和j的元素位置
	sort.Ints(nums[i+1:])               // 后半段升序排列
	copy(res[i:], nums[i:])
	for _, v := range res {
		ans = ans*10 + v
	}
	if ans > math.MaxInt32 {
		return -1
	}
	return
}
