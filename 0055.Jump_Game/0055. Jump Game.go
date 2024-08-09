package main

func canJump(nums []int) bool {
	var mx int // 维护能跳到的最右位置
	for i, jump := range nums {
		if i > mx {
			return false
		}
		mx = max(mx, i+jump)
	}
	return true
}
