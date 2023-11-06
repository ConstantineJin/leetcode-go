package main

// 从最后一间房屋开始打劫
//func rob(nums []int) int {
//	var mem = make([]int, len(nums)) // mem数组保存该房屋的dfs结果，若dfs[i]==-1表明第i号房屋未计算过
//	for i := range mem {             // 将mem的全部元素初始化为-1
//		mem[i] = -1
//	}
//	var dfs func(i int) int
//	dfs = func(i int) int {
//		if i < 0 { // 已经打劫到头了，返回0
//			return 0
//		}
//		if mem[i] != -1 { //计算过结果，直接读取返回
//			return mem[i]
//		}
//		mem[i] = max(nums[i]+dfs(i-2), dfs(i-1)) // 打劫该房屋收益：nums[i]+dfs(i-2)，不打劫该房屋收益：dfs(i-1)。选取两者的最大值并保存结果
//		return mem[i]
//	}
//	return dfs(len(nums) - 1) // 从最后一间房屋开始打劫
//}

// 从头开始打劫
//func rob(nums []int) int {
//	mem := make([]int, len(nums))
//	for i := range mem {
//		mem[i] = -1
//	}
//	var dfs func(i int) int
//	dfs = func(i int) int {
//		if i >= len(nums) {
//			return 0
//		}
//		if mem[i] != -1 {
//			return mem[i]
//		}
//		mem[i] = max(nums[i]+dfs(i+2), dfs(i+1))
//		return mem[i]
//	}
//	return dfs(0)
//}

// 改递归为递推
//func rob(nums []int) int {
//	n := len(nums)
//	f := make([]int, n+2) // 为避免第0和第1个需要特殊讨论,直接全体后移2个
//	for i, num := range nums {
//		f[i+2] = max(f[i]+num, f[i+1])
//	}
//	return f[n+1]
//}

// 递推空间复杂度从O(n)优化到O(1)
func rob(nums []int) int {
	x, y := 0, 0
	for _, num := range nums {
		z := max(x+num, y)
		x, y = y, z
	}
	return y
}
