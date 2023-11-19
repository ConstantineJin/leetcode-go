package main

// 滑动窗口
func maxSumOfThreeSubarrays(nums []int, k int) []int {
	var sum1, sum2, sum3, maxSum1, maxSum12, maxSum123, maxSum1Idx, maxSum12Idx1, maxSum12Idx2 int // sum记录每个窗口当前的和，maxSum记录前i个窗口取到过的最大和，maxSumIdx记录对应窗口的起始元素下标
	ans := make([]int, 3)
	for i := k * 2; i < len(nums); i++ { // 保证三个窗口没有重叠部分
		sum1 += nums[i-k*2] // 窗口右端后移，添加新元素
		sum2 += nums[i-k]
		sum3 += nums[i]
		if i >= k*3-1 {
			if sum1 > maxSum1 { // 按需更新最大值及下标
				maxSum1, maxSum1Idx = sum1, i-k*3+1
			}
			if maxSum1+sum2 > maxSum12 {
				maxSum12, maxSum12Idx1, maxSum12Idx2 = maxSum1+sum2, maxSum1Idx, i-k*2+1
			}
			if maxSum12+sum3 > maxSum123 {
				maxSum123, ans[0], ans[1], ans[2] = maxSum12+sum3, maxSum12Idx1, maxSum12Idx2, i-k+1
			}
			sum1 -= nums[i-k*3+1] // 窗口前端后移，移除旧元素
			sum2 -= nums[i-k*2+1]
			sum3 -= nums[i-k+1]
		}
	}
	return ans
}
