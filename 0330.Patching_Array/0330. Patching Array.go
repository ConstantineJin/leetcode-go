package main

func minPatches(nums []int, n int) (ans int) {
	s, i := 1, 0
	for s <= n {
		if i < len(nums) && nums[i] <= s {
			s += nums[i]
			i++
		} else {
			s *= 2 // 必须添加 s
			ans++
		}
	}
	return
}
