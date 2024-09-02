package main

func maxScore(grid [][]int) int {
	pos := make(map[int]int) // key 是元素值，value 是压缩为二进制数的元素出现的行号
	for i, row := range grid {
		for _, v := range row {
			pos[v] |= 1 << i
		}
	}
	allNums := make([]int, 0, len(pos))
	for x := range pos {
		allNums = append(allNums, x)
	}
	f := make([][]int, len(allNums)+1) // f[i][j] 表示在 [1,i] 中选数字，已选择行号集合为 j（即 j 的每一个为 1 的比特位对应的行号不可选择），每行至多一个数且没有相同元素时所选元素之和的最大值
	for i := range f {
		f[i] = make([]int, 1<<len(grid))
	}
	for i, x := range allNums {
		for j, v := range f[i] {
			f[i+1][j] = v                                    // 不选 x
			for t, lowBit := pos[x], 0; t > 0; t ^= lowBit { // 选 x，遍历 x 出现的每一行
				lowBit = t & -t    // lowBit = 1<<k，其中 k 是行号
				if j&lowBit == 0 { // 没选过第 k 行的数
					f[i+1][j] = max(f[i+1][j], f[i][j|lowBit]+x) // 选第 k 行的 x
				}
			}
		}
	}
	return f[len(allNums)][0]
}
