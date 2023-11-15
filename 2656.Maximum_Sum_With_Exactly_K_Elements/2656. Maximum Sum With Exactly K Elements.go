package main

func maximizeSum(nums []int, k int) int {
	m := nums[0]
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	return m*k + (k-1)*k/2 // 贪心，每次选择数组中的最大值，+1后仍是最大值
}
