package main

func intersection(nums1 []int, nums2 []int) (ans []int) {
	var cnt1, cnt2 [1001]bool
	for _, num := range nums1 {
		cnt1[num] = true
	}
	for _, num := range nums2 {
		cnt2[num] = true
	}
	for i := range cnt1 {
		if cnt1[i] && cnt2[i] {
			ans = append(ans, i)
		}
	}
	return
}
