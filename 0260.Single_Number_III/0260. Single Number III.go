package main

func singleNumber(nums []int) []int {
	var xorSum int
	// 第一遍遍历获取全体元素的异或和，亦即两个只出现一次的元素的异或和
	for _, num := range nums {
		xorSum ^= num
	}
	// lowBit为二进制为1的最低位，通过n&-n求得（可以取任一为1的二进制位，并不一定要取lowBit）
	lowBit, ans := xorSum&-xorSum, make([]int, 2)
	for _, num := range nums {
		// 对于该比特位，两个只出现一次的数在该位一个是0一个是1，必然在不同组，此后解法同0136
		if num&lowBit == 0 {
			ans[0] ^= num
		} else {
			ans[1] ^= num
		}
	}
	return ans
}
