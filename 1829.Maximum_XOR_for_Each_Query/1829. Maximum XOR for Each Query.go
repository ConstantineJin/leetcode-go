package main

func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	mask := (1 << maximumBit) - 1
	ans := make([]int, n)
	var xorSum int
	for i, num := range nums {
		xorSum ^= num
		ans[n-i-1] = mask ^ xorSum
	}
	return ans
}
