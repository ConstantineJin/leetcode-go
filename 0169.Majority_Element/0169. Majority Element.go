package main

// 哈希表，空间复杂度O(n)
//func majorityElement(nums []int) int {
//	mp := make(map[int]int)
//	for _, num := range nums {
//		mp[num]++
//	}
//	for i, v := range mp {
//		if v > len(nums)/2 {
//			return i
//		}
//	}
//	return -1
//}

// 摩尔投票法，空间复杂度O(1)
func majorityElement(nums []int) int {
	major, vote := nums[0], 0  // major记录当前得票最高的元素，vote记录其票数
	for _, num := range nums { // 依次遍历
		if vote == 0 { // 如果major元素当前得票为0
			major = num //那么将major的值更新为当前遍历到的元素
			vote++
			continue
		}
		if num == major { // 当前遍历到的元素值为major
			vote++ // 票数+1
		} else {
			vote-- // 否则票数-1
		}
	}
	return major
}
