package main

// 方法1: 枚举 X
//func minChanges(nums []int, k int) int {
//	n := len(nums)
//	cnt, cnt2 := make([]int, k+1), make([]int, k+1) // cnt 维护 q-p 的出现次数，cnt2 维护 max(q, k-p) 的出现次数
//	for i, p := range nums[:n/2] {
//		// p, q = min(nums[i], nums[n-i-1]), max(nums[i], nums[n-i-1])
//		// 如果 q-p = X，不需要修改
//		// 如果 q-p > X，只需修改一个数
//		// 如果 q-p < X 且 max(q, k-p) >= X，只需修改一个数
//		// 如果 q-p < X 且 max(q, k-p) < X，需要修改两个数
//		q := nums[n-i-1]
//		if p > q {
//			p, q = q, p
//		}
//		cnt[q-p]++
//		cnt2[max(q, k-p)]++
//	}
//	ans := n                // 极端情况：全部都要修改
//	var sum2 int            // cnt2 的前缀和
//	for x, c := range cnt { // X 的取值在 [0,k] 区间内
//		ans = min(ans, n/2-c+sum2) // 有 cnt[X] 对不需要修改，有 n-cnt[X] 对需要修改至少一个，有 cnt2[0]+cnt2[1]+...+cnt2[X-1] 对还需要再修改一个
//		sum2 += cnt2[x]
//	}
//	return ans
//}

// 方法2: 差分数组
func minChanges(nums []int, k int) int {
	n := len(nums)
	diff := make([]int, k+2)
	for i, p := range nums[:n/2] {
		// p, q = min(nums[i], nums[n-i-1]), max(nums[i], nums[n-i-1])
		// 如果 q-p = X，不需要修改
		// 如果 q-p > X，只需修改一个数
		// 如果 q-p < X 且 max(q, k-p) >= X，只需修改一个数
		// 如果 q-p < X 且 max(q, k-p) < X，需要修改两个数
		q := nums[n-i-1]
		if p > q {
			p, q = q, p
		}
		x := q - p
		mx := max(q, k-p)
		// [0, x-1] 全部 +1：把 q-p 改成小于 x 的，只需要改 p 或 q 中的一个数
		diff[0]++
		diff[x]--
		// [x+1, mx] 全部 +1：把 q-p 改成大于 x 小于等于 mx 的，也只需要改 p 或 q 中的一个数
		diff[x+1]++
		// [mx+1, k] 全部 +2：把 q-p 改成大于 mx 的，p 和 q 都需要改
		diff[mx+1]++
	}
	ans := n // 极端情况：全部都要修改
	var mn int
	for _, d := range diff {
		mn += d
		ans = min(ans, mn)
	}
	return ans
}
