package main

import "sort"

// 方法1: 暴力+快速选择
//func quickSelect(nums []int32, l, r, k int) int {
//	if l == r {
//		return int(nums[l])
//	}
//	partition := nums[l]
//	i, j := l-1, r+1
//	for i < j {
//		for i++; i < r && nums[i] < partition; i++ {
//		}
//		for j--; j > l && nums[j] > partition; j-- {
//		}
//		if i < j {
//			nums[i], nums[j] = nums[j], nums[i]
//		}
//	}
//	if k <= j {
//		return quickSelect(nums, l, j, k)
//	} else {
//		return quickSelect(nums, j+1, r, k)
//	}
//}
//
//func smallestDistancePair(nums []int, k int) int {
//	sort.Ints(nums)
//	n := len(nums)
//	distances := make([]int32, 0, n*(n-1)/2)
//	for i := 0; i < n; i++ {
//		for j := i + 1; j < n; j++ {
//			distances = append(distances, int32(nums[j]-nums[i]))
//		}
//	}
//	return quickSelect(distances, 0, len(distances)-1, k-1)
//}

// 方法2: 排序+二分查找
//func smallestDistancePair(nums []int, k int) int {
//	sort.Ints(nums)
//	return sort.Search(nums[len(nums)-1]-nums[0], func(mid int) bool { // 二分查找距离 mid
//		var cnt int
//		for i, num := range nums {
//			j := sort.SearchInts(nums[:i], num-mid) // 二分查找大于等于 nums[i]−mid 的最小值的下标 j
//			cnt += i - j                            // 右端点为 i 且距离小于等于 mid 的数对数目为 i-j
//		}
//		return cnt >= k
//	})
//}

// 方法3: 排序+双指针/滑动窗口
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	return sort.Search(nums[len(nums)-1]-nums[0], func(mid int) bool { // 二分查找距离 mid
		var cnt, i int
		for j, num := range nums {
			for num-nums[i] > mid {
				i++
			}
			cnt += j - i
		}
		return cnt >= k
	})
}
