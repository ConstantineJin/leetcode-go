package main

func numSpecial(mat [][]int) (ans int) {
	var m, n = len(mat), len(mat[0])
	var rowSum, colSum = make([]int, m), make([]int, n)
	for i, row := range mat {
		for j, v := range row {
			rowSum[i] += v
			colSum[j] += v
		}
	}
	for i, row := range mat {
		if rowSum[i] != 1 {
			continue
		}
		for j, v := range row {
			if colSum[j] != 1 {
				continue
			}
			if v == 1 {
				ans++
			}
		}
	}
	return
}
