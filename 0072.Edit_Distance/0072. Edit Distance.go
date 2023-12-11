package main

// 递归
//func minDistance(word1 string, word2 string) int {
//	var m, n = len(word1), len(word2)
//	mem := make([][]int, m) // 记忆化搜索
//	for i := range mem {
//		mem[i] = make([]int, n)
//		for j := range mem[i] {
//			mem[i][j] = -1
//		}
//	}
//	var dfs func(int, int) int
//	dfs = func(i int, j int) (res int) {
//		if i < 0 { // 如果i为负
//			return j + 1 // 说明前面还需要添加j+1个字母
//		}
//		if j < 0 { // 如果j为负
//			return i + 1 // 说明前面还需要删除i+1个字母
//		}
//		var p = &mem[i][j]
//		if *p != -1 { // 如果已经计算过了
//			return *p
//		}
//		defer func() { *p = res }() // 保存计算结果
//		if word1[i] == word2[j] {   // 当前两个字符相同
//			return dfs(i-1, j-1) // 不需要操作
//		} else {
//			return min(dfs(i-1, j), dfs(i, j-1), dfs(i-1, j-1)) + 1 // 分别对应删除，插入和替换
//		}
//	}
//	return dfs(m-1, n-1)
//}

// 递推
//func minDistance(word1 string, word2 string) int {
//	var m, n = len(word1), len(word2)
//	var f = make([][]int, m+1)
//	for i := range f {
//		f[i] = make([]int, n+1)
//	}
//	for j := 1; j <= n; j++ { // 边界初始化
//		f[0][j] = j
//	}
//	for i, x := range word1 {
//		f[i+1][0] = i + 1 // 边界初始化
//		for j, y := range word2 {
//			if x == y { // 如果当前两个字符相等
//				f[i+1][j+1] = f[i][j] // 不需要编辑
//			} else { // 如果当前两个字符不相等
//				f[i+1][j+1] = min(f[i][j+1], f[i+1][j], f[i][j]) + 1 // 选择添加、删除和替换三种方案中的最小值
//			}
//		}
//	}
//	return f[m][n]
//}

// 递推空间优化，使用两个一维切片
func minDistance(word1 string, word2 string) int {
	var m, n = len(word1), len(word2)
	var f = [2][]int{make([]int, n+1), make([]int, n+1)}
	for j := 1; j <= n; j++ { // 边界初始化
		f[0][j] = j
	}
	for i, x := range word1 {
		f[(i+1)%2][0] = i + 1
		for j, y := range word2 {
			if x == y {
				f[(i+1)%2][j+1] = f[i%2][j]
			} else {
				f[(i+1)%2][j+1] = min(f[i%2][j+1], f[(i+1)%2][j], f[i%2][j]) + 1
			}
		}
	}
	return f[m%2][n]
}
