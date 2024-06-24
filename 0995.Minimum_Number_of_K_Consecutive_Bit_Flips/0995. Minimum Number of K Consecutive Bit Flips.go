package main

func minKBitFlips(nums []int, k int) (ans int) {
	n := len(nums)
	diff := make([]int, n+1) // 差分数组，diff[i] 表示两个相邻元素 nums[i−1] 和 nums[i] 的翻转次数的差
	var cnt int              // 当前翻转次数计数
	for i, num := range nums {
		cnt ^= diff[i]
		if num == cnt { // (num + cnt) % 2 == 0，需要从当前元素开始翻转
			if i+k > n { // 剩余元素数量不足以完成一次翻转
				return -1
			}
			ans++
			cnt ^= 1
			diff[i+k] ^= 1
		}
	}
	return
}
