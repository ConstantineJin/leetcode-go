package main

func minOperations(nums []int, k int) (ans int) {
	for _, num := range nums {
		if num < k {
			ans++
		}
	}
	return ans
}
