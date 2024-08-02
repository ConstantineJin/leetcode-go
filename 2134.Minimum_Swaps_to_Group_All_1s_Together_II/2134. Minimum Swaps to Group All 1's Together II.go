package main

func minSwaps(nums []int) int {
	var tot1, cnt1, maxCnt1 int
	for _, num := range nums {
		tot1 += num
	}
	nums = append(nums, nums...) // 环状数组
	for i, num := range nums {
		cnt1 += num
		if i >= tot1 { // 滑动窗口，大小固定为 tot1
			cnt1 -= nums[i-tot1]
			maxCnt1 = max(maxCnt1, cnt1)
		}
	}
	return tot1 - maxCnt1 // 滑动窗口中 0 的最小数量就是需要的最小交换次数
}
