package main

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// 对于一个长度为n的数组，其中没有出现的最小正整数只能在[1,n+1]中
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 { // 对于0和负数，将其设为n+1
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num < n+1 {
			nums[num-1] = -abs(nums[num-1]) // 通过设置对应下标的数为负来实现不需要额外空间的标记
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1 // 结果是遍历后第一个正数的下标+1
		}
	}
	return n + 1 // 如果全是负数或0，那么结果就是n+1
}
