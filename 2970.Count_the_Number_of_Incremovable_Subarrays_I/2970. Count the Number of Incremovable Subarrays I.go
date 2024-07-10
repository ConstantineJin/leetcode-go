package main

func incremovableSubarrayCount(nums []int) (ans int) {
	n := len(nums)
	left, right := 0, n-1 // left 和 right 分别指向以 nums[0] 开始和以 nums[n-1] 结束的严格递增子数组的另一边界
	for left < n-1 && nums[left+1] > nums[left] {
		left++
	}
	for right > 0 && nums[right-1] < nums[right] {
		right--
	}
	if left >= right { // nums 本身严格单调
		return (n + 1) * n / 2 // nums 的任一非空子数组都是移除递增子数组
	}
	ans = n - right + 1 // 不取左半边，则从 nuns[right:n] 的任一子数组都满足要求
	j := right
	for i := 0; i <= left; i++ { // 双指针
		for j < n && nums[j] <= nums[i] { // nums[j] 必须严格大于 nums[i]
			j++
		}
		ans += n - j + 1 // 对于当前的 i，nums[j:n] 的任一子数组都满足要求
	}
	return
}
