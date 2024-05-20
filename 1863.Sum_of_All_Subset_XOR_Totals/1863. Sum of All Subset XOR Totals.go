package main

// 引理1:如果nums中任何一个数在这一位上是0，那么任何一个子集的异或结果在这一位上都是0
// 引理2:如果nums中有一个数字在这一位上是1，那么所有子集异或结果在这一位上一半是0，一半是1（使用数学归纳法证明）
func subsetXORSum(nums []int) (ans int) {
	for _, num := range nums {
		ans |= num
	}
	return ans << (len(nums) - 1)
}
