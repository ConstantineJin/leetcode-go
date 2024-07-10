package main

func countSubarrays(nums []int, k int) (ans int64) {
	var left, right int // [left,right) 是满足 nums[j]=k 的 j 的范围
	for i, num := range nums {
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
