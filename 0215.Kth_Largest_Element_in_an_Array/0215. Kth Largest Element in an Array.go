package main

// 全排序
//func findKthLargest(nums []int, k int) int {
//	sort.Ints(nums)
//	return nums[len(nums)-k]
//}

// 快速选择算法
func quickSelect(s []int, left, right, k int) int {
	if left == right {
		return s[left]
	}
	pivot := s[left]
	i, j := left-1, right+1
	for i < j {
		for i++; s[i] < pivot; i++ {
		}
		for j--; s[j] > pivot; j-- {
		}
		if i < j {
			s[i], s[j] = s[j], s[i]
		}
	}
	if k <= j {
		return quickSelect(s, left, j, k)
	} else {
		return quickSelect(s, j+1, right, k)
	}
}

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}
