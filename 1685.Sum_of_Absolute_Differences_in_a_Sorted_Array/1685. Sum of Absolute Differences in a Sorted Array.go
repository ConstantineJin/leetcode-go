package main

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	var s int
	for _, num := range nums {
		s += num
	}
	ans[0] = s - nums[0]*n
	for i := 1; i < n; i++ {
		diff := nums[i] - nums[i-1]
		ans[i] = ans[i-1] + diff*i - diff*(n-i)
	}
	return ans
}
