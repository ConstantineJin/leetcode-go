package main

func missingNumber(nums []int) int {
	var n = len(nums)
	var sum int
	for _, num := range nums {
		sum += num
	}
	return n*(n+1)/2 - sum
}
