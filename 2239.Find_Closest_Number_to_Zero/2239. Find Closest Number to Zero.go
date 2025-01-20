package main

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findClosestNumber(nums []int) int {
	ans := nums[0]
	for _, num := range nums[1:] {
		if abs(num) < abs(ans) || abs(num) == abs(ans) && num > 0 {
			ans = num
		}
	}
	return ans
}
