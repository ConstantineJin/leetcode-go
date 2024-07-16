package main

func findIntersectionValues(nums1, nums2 []int) []int {
	mp1, mp2, ans := make(map[int]int), make(map[int]int), make([]int, 2)
	for _, num := range nums1 {
		mp1[num] = 1
	}
	for _, num := range nums2 {
		mp2[num] = 1
	}
	for _, num := range nums1 {
		ans[0] += mp2[num]
	}
	for _, num := range nums2 {
		ans[1] += mp1[num]
	}
	return ans
}
