package main

// 方法1:记忆化搜索
//func maxCoins(nums []int) (ans int) {
//	n := len(nums)
//	nums = append(append([]int{1}, nums...), 1)
//	mem := make([][]int, n+2)
//	for i := range mem {
//		mem[i] = make([]int, n+2)
//		for j := range mem[i] {
//			mem[i][j] = -1
//		}
//	}
//	var dfs func(left, right int) int // 反向思考。dfs(left, right)表示将开区间(left, right)内全部填满气球能够得到的最多硬币数
//	dfs = func(left, right int) int {
//		if left >= right-1 {
//			return 0
//		}
//		if mem[left][right] == -1 {
//			for i := left + 1; i < right; i++ { // 枚举每一个中间位置
//				sum := nums[left] * nums[i] * nums[right]
//				sum += dfs(left, i) + dfs(i, right)
//				mem[left][right] = max(mem[left][right], sum)
//			}
//		}
//		return mem[left][right]
//	}
//	return dfs(0, n+1)
//}

// 方法2:递推
func maxCoins(nums []int) (ans int) {
	n := len(nums)
	nums = append(append([]int{1}, nums...), 1)
	dp := make([][]int, n+2) // dp[i][j]表示填满开区间(i,j)能得到的最多硬币数
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				sum := nums[i] * nums[k] * nums[j]
				sum += dp[i][k] + dp[k][j]
				dp[i][j] = max(dp[i][j], sum)
			}
		}
	}
	return dp[0][n+1]
}
