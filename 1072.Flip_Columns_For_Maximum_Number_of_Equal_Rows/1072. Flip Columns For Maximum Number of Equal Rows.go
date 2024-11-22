package main

func maxEqualRowsAfterFlips(matrix [][]int) (ans int) {
	mp := make(map[[5]uint]int)
	for _, row := range matrix {
		var r [5]uint
		for i, v := range row {
			r[i/64] |= uint(v^row[0]) << (i % 64) // 如果该列第一行为 1，那么翻转该列
		}
		mp[r]++
	}
	for _, v := range mp {
		ans = max(ans, v)
	}
	return
}
