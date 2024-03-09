package main

func getCommon(nums1 []int, nums2 []int) int {
	var m, n = len(nums1), len(nums2)
	var i, j int
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			return nums1[i]
		}
	}
	return -1
}
