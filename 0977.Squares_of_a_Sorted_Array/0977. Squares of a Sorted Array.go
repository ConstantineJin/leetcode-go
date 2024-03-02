package main

func sortedSquares(nums []int) []int {
	var n = len(nums)
	var ans, i, j = make([]int, n), 0, n - 1
	for k := n - 1; i <= j; k-- {
		var x, y = nums[i] * nums[i], nums[j] * nums[j]
		if x < y {
			ans[k] = y
			j--
		} else {
			ans[k] = x
			i++
		}
	}
	return ans
}
