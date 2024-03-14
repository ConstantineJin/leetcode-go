package main

// 前缀和+哈希表
//func numSubarraysWithSum(nums []int, goal int) (ans int) {
//	var mp = make(map[int]int)
//	var sum int
//	for _, num := range nums {
//		mp[sum]++
//		sum += num
//		ans += mp[sum-goal]
//	}
//	return
//}

// 前缀和+数组
func numSubarraysWithSum(nums []int, goal int) (ans int) {
	var n = len(nums)
	var preSum = make([]int, n+1)
	var arr = make([]int, 1) // 下标：前缀和；值：个数
	arr[0], preSum[0] = 1, 0
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
		if nums[i-1] == 1 {
			arr = append(arr, 1)
		} else {
			arr[len(arr)-1]++
		}
	}
	if goal == 0 {
		for _, v := range arr {
			ans += v * (v - 1) / 2
		}
	} else {
		for i := 0; i+goal < len(arr); i++ {
			ans += arr[i] * arr[i+goal]
		}
	}
	return
}
