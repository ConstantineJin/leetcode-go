package main

//func alternatingSubarray(nums []int) (ans int) {
//	var cnt, diff int
//	for i := 1; i < len(nums); i++ {
//		switch nums[i] - nums[i-1] {
//		case 1:
//			if diff == -1 && cnt > 1 {
//				cnt++
//			} else {
//				cnt = 2
//			}
//			diff = 1
//		case -1:
//			if diff == 1 {
//				cnt++
//			} else {
//				cnt = 0
//			}
//			diff = -1
//		default:
//			cnt, diff = 0, 0
//		}
//		ans = max(ans, cnt)
//	}
//	if ans < 2 {
//		return -1
//	}
//	return
//}

func alternatingSubarray(nums []int) int {
	var ans, i, n = -1, 0, len(nums)
	for i < n-1 {
		if nums[i+1]-nums[i] != 1 {
			i++ // 直接跳过
			continue
		}
		var i0 = i // 记录这一组的开始位置
		i += 2     // i 和 i+1 已经满足要求，从 i+2 开始判断
		for i < n && nums[i] == nums[i0]+(i-i0)%2 {
			i++
		}
		// 从 i0 到 i-1 是满足题目要求的（并且无法再延长的）子数组
		ans = max(ans, i-i0)
		i--
	}
	return ans
}
