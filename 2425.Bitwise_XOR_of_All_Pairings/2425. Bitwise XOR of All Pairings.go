package main

func xorAllNums(nums1, nums2 []int) (ans int) {
	if len(nums1)&1 > 0 {
		for _, num := range nums2 {
			ans ^= num
		}
	}
	if len(nums2)&1 > 0 {
		for _, num := range nums1 {
			ans ^= num
		}
	}
	return
}
