package main

// 动态规划
func findNumberOfLIS(nums []int) (ans int) {
	var n, maxLen = len(nums), 0                 // maxLen维护最长上升子序列的长度
	var dp, cnt = make([]int, n), make([]int, n) // dp[i]表示以nums[i]结尾的LIS的长度，cnt[i]表示以nums[i]结尾的LIS的个数
	for i := range nums {
		dp[i], cnt[i] = 1, 1
		for j := range nums[:i] {
			if nums[i] > nums[j] { // 当前遍历到的nums[i]之前有更大的元素nums[j]
				if dp[i] < dp[j]+1 { // 如果以nums[j]结尾的LIS在加上nums[i]之后长度比先前记录的以nums[i]结尾的LIS更长
					dp[i], cnt[i] = dp[j]+1, cnt[j] // 更新以nums[i]结尾的LIS长度和以nums[i]结尾的LIS的个数
				} else if dp[i] == dp[j]+1 { // 如果以nums[j]结尾的LIS在加上nums[i]之后长度比先前记录的以nums[i]结尾的LIS等长
					cnt[i] += cnt[j] // 新的以nums[i]结尾的LIS的个数等于以nums[i]结尾的LIS的个数再加上以nums[j]结尾的LIS的个数
				}
			}
		}
		if dp[i] > maxLen { // 外层循环更新同内层循环
			maxLen, ans = dp[i], cnt[i]
		} else if dp[i] == maxLen {
			ans += cnt[i]
		}
	}
	return
}
