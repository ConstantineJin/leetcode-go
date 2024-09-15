package main

// 方法1: 排序，同 0065.合并区间
//func numberOfPoints(nums [][]int) (ans int) {
//	slices.SortFunc(nums, func(a, b []int) int { return a[0] - b[0] })
//	left, right := nums[0][0], nums[0][1]
//	for i := 1; i < len(nums); i++ {
//		if nums[i][0] <= right {
//			right = max(right, nums[i][1])
//		} else {
//			ans += right - left + 1
//			left, right = nums[i][0], nums[i][1]
//		}
//	}
//	ans += right - left + 1
//	return
//}

// 方法2: 差分数组
func numberOfPoints(nums [][]int) (ans int) {
	var maxEnd int
	for _, num := range nums {
		maxEnd = max(maxEnd, num[1])
	}
	diff := make([]int, maxEnd+2)
	for _, interval := range nums {
		diff[interval[0]]++
		diff[interval[1]+1]--
	}
	var s int
	for _, d := range diff {
		s += d
		if s > 0 {
			ans++
		}
	}
	return
}
