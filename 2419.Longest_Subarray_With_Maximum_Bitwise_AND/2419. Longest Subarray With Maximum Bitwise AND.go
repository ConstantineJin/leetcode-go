package main

func longestSubarray(nums []int) (ans int) { // 两数按位与的结果不可能比两数中的任何一数更大，因此本题等价于求数组中最大值连续出现的最长长度
	var mx, cnt int
	for _, num := range nums {
		if num > mx {
			mx, ans, cnt = num, 1, 1
		} else if num < mx {
			cnt = 0
		} else if cnt++; cnt > ans {
			ans = cnt
		}
	}
	return ans
}
