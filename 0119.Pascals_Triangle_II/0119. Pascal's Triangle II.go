package main

// 方法1: 模拟
//func getRow(rowIndex int) []int {
//	ans := make([]int, 1, rowIndex+1)
//	ans[0] = 1
//	for i := 0; i < rowIndex; i++ {
//		ans = append(ans, 1)
//		for j := i; j > 0; j-- {
//			ans[j] += ans[j-1]
//		}
//	}
//	return ans
//}

// 方法2: 数学+线性递推
func getRow(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	ans[0] = 1
	for i := 1; i <= rowIndex; i++ {
		ans[i] = ans[i-1] * (rowIndex - i + 1) / i
	}
	return ans
}
