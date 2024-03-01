package main

func validPartition(nums []int) bool {
	var n = len(nums)
	var f = make([]bool, n+1)
	f[0] = true
	for i, num := range nums {
		if i > 0 && f[i-1] && num == nums[i-1] || i > 1 && f[i-2] && (num == nums[i-1] && num == nums[i-2] || num == nums[i-1]+1 && num == nums[i-2]+2) {
			f[i+1] = true
		}
	}
	return f[n]
}
