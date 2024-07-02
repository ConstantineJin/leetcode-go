package main

// 方法1: 排序
//func intersect(nums1 []int, nums2 []int) (ans []int) {
//	sort.Ints(nums1)
//	sort.Ints(nums2)
//	m, n := len(nums1), len(nums2)
//	var i, j int
//	for i < m && j < n {
//		if nums1[i] < nums2[j] {
//			i++
//		} else if nums1[i] > nums2[j] {
//			j++
//		} else {
//			ans = append(ans, nums1[i])
//			i++
//			j++
//		}
//	}
//	return
//}

// 方法2: 哈希表
func intersect(nums1 []int, nums2 []int) (ans []int) {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	mp := make(map[int]int)
	for _, num := range nums1 {
		mp[num]++
	}
	for _, num := range nums2 {
		if mp[num] > 0 {
			ans = append(ans, num)
			mp[num]--
		}
	}
	return
}
