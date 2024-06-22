package main

func numberOfSubarrays(nums []int, k int) (ans int) {
	oddIdx := []int{-1} // 所有奇数的下标，在首尾分别添加 -1 和 len(nums) 方便边界处理
	for i, num := range nums {
		if num%2 == 1 {
			oddIdx = append(oddIdx, i)
		}
	}
	oddIdx = append(oddIdx, len(nums))
	for i := 1; i < len(oddIdx)-k; i++ { // 滑动窗口
		ans += (oddIdx[i] - oddIdx[i-1]) * (oddIdx[i+k] - oddIdx[i+k-1]) // 乘法原理
	}
	return
}
