package main

func minBitwiseArray(nums []int) []int {
	for i, num := range nums {
		if num == 2 {
			nums[i] = -1
		} else { // x|(x+1) 的本质是把 x 的二进制表示中最右侧的 0 置 1
			nums[i] ^= (num + 1) &^ num >> 1
		}
	}
	return nums
}
