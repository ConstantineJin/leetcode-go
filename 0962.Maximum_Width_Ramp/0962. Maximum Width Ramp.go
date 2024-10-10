package main

func maxWidthRamp(nums []int) (ans int) {
	var st []int
	for i, num := range nums {
		for len(st) == 0 || num < nums[st[len(st)-1]] {
			st = append(st, i)
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(st) > 0 && nums[st[len(st)-1]] <= nums[i] {
			ans = max(ans, i-st[len(st)-1])
			st = st[:len(st)-1]
		}
	}
	return
}
