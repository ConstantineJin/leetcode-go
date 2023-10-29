package main

// 排序
//func hIndex(citations []int) (h int) {
//	sort.Ints(citations)
//	for i := len(citations); i > 0 && citations[i-1] > h; i-- {
//		h++
//	}
//	return
//}

// 计数排序
//func min(a, b int) int {
//	if a < b {
//		return a
//	} else {
//		return b
//	}
//}
//
//func hIndex(citations []int) (h int) {
//	n := len(citations)
//	cnt := make([]int, n+1)
//	for _, citation := range citations {
//		cnt[min(citation, n)]++
//	}
//	for i := n; i > 0; i-- {
//		h += cnt[i]
//		if h >= i {
//			return i
//		}
//	}
//	return
//}

// 二分查找
func hIndex(citations []int) int {
	left, right := 0, len(citations)
	for left < right {
		mid := (left + right + 1) >> 1
		cnt := 0
		for _, v := range citations {
			if v >= mid {
				cnt++
			}
		}
		if cnt >= mid {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}
