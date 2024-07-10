package main

func countSubarrays(nums []int, k int) (ans int64) {
	var left, right int        // [left,right) 是满足 nums[j]=k 的 j 的范围
	for i, num := range nums { // 正向遍历 nums
		// 对于 num=nums[i]，从 i−1 开始倒着遍历 nums[j]，更新 nums[j]=nums[j]&num。
		// 如果 nums[j]&num=nums[j]，说明 num 对应的集合是 nums[j] 的超集，也是 nums[k] (0<=k<=j) 的超集，后续遍历不可能 AND 值，跳出内循环
		for j := i - 1; j >= 0 && nums[j]&num != nums[j]; j-- {
			nums[j] &= num
		}
		for left <= i && nums[left] < k {
			left++
		}
		for right <= i && nums[right] <= k {
			right++
		}
		ans += int64(right - left)
	}
	return
}
