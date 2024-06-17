package main

func minPatches(nums []int, n int) (ans int) {
	s := 1 // [0, s)区间内的都已经得到
	for i := 0; s <= n; {
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
