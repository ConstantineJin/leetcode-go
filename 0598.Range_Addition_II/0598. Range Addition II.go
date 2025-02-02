package main

func maxCount(m, n int, ops [][]int) int {
	minX, minY := m, n
	for _, op := range ops {
		minX, minY = min(minX, op[0]), min(minY, op[1])
	}
	return minX * minY
}
