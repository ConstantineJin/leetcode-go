package main

// 二进制枚举
//func maximumRows(mat [][]int, numSelect int) (ans int) {
//	var m, n = len(mat), len(mat[0])
//	var mask = make([]int, m)
//	for i, row := range mat {
//		for j, x := range row {
//			mask[i] |= x << j // 初始化mask，将每行的01序列转化为二进制数
//		}
//	}
//	for subset := 0; subset < 1<<n; subset++ {
//		if bits.OnesCount(uint(subset)) != numSelect { // 如果当前枚举到的子集中1的数量不等于numSelect
//			continue // 继续枚举下一个子集
//		}
//		var coveredRows int
//		for _, row := range mask {
//			if row&subset == row { // 如果当前行与当前枚举到的子集等于自身，说明当前行覆盖了当前子集
//				coveredRows++
//			}
//		}
//		ans = max(ans, coveredRows)
//	}
//	return
//}

// Gosper's Hack
func maximumRows(mat [][]int, numSelect int) (ans int) {
	var m, n = len(mat), len(mat[0])
	var mask = make([]int, m)
	for i, row := range mat {
		for j, x := range row {
			mask[i] |= x << j // 初始化mask，将每行的01序列转化为二进制数
		}
	}
	var subset = 1<<numSelect - 1 // 初始化subset为第一个k元子集（00...0011...11（k个1））
	for subset < 1<<n {
		var coveredRows int
		for _, row := range mask {
			if row&subset == row { // 如果当前行与当前枚举到的子集等于自身，说明当前行覆盖了当前子集
				coveredRows++
			}
		}
		ans = max(ans, coveredRows)
		var lowBit = subset & -subset
		var x = subset + lowBit
		subset = (subset^x)/lowBit>>2 | x // 无需枚举所有子集，只需枚举下一个k元子集
	}
	return
}
