package main

func goodSubsetofBinaryMatrix(grid [][]int) []int {
	maskToIdx := make(map[int]int) // key 为二进制数，value为行号
	for i, row := range grid {
		var mask int
		for j, x := range row {
			mask |= x << j
		}
		if mask == 0 { // 如果有一行全为0，直接返回
			return []int{i}
		}
		maskToIdx[mask] = i
	}
	for x, i := range maskToIdx {
		for y, j := range maskToIdx {
			if x&y == 0 {
				return []int{min(i, j), max(i, j)}
			}
		}
	}
	return nil
}
