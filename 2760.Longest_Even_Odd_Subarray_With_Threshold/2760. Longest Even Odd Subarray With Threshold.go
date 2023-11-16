package main

func longestAlternatingSubarray(nums []int, threshold int) (ans int) {
	n, i, j := len(nums), 0, 0
	for i < n {
		if nums[i]%2 != 0 || nums[i] > threshold { // 检查当前子数组的起始元素是否满足①是2的倍数②是否不超过threshold
			i++ // 如果不满足就检查下一个
			continue
		}
		for j = i + 1; j < n && nums[j] <= threshold && nums[j]%2 != nums[j-1]%2; j++ {
		} // 从当前子数组起始元素的下一个元素开始检查，如果不满足①元素小于等于threshold②与前一个元素的奇偶性不同，那么前一个元素就是子数组的结束
		ans, i = max(ans, j-i), j
	}
	return
}
