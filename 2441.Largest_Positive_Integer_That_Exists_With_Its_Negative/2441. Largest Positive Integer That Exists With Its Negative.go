package main

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findMaxK(nums []int) int {
	var has, ans = make(map[int]bool), -1
	for _, num := range nums {
		if has[-num] && abs(num) > ans {
			ans = abs(num)
		}
		has[num] = true
	}
	return ans
}
