package main

func getKthElement(nums1, nums2 []int, k int) int { // 寻找两个有序数组中第k大的元素
	var i, j int
	for {
		if i == len(nums1) {
			return nums2[j+k-1]
		}
		if j == len(nums2) {
			return nums1[i+k-1]
		}
		if k == 1 {
			return min(nums1[i], nums2[j])
		}
		var ii, jj = min(i+k/2, len(nums1)) - 1, min(j+k/2, len(nums2)) - 1
		if nums1[ii] <= nums2[jj] {
			k -= ii - i + 1
			i = ii + 1
		} else {
			k -= jj - j + 1
			j = jj + 1
		}
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var n = len(nums1) + len(nums2)
	if n%2 == 1 {
		return float64(getKthElement(nums1, nums2, n/2+1))
	} else {
		return float64(getKthElement(nums1, nums2, n/2)+getKthElement(nums1, nums2, n/2+1)) / 2
	}
}
