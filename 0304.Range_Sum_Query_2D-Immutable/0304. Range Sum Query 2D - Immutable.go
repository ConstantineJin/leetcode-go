package main

type NumMatrix struct{ preSum [][]int }

func Constructor(matrix [][]int) NumMatrix {
	var m, n = len(matrix), len(matrix[0])
	var numMatrix = NumMatrix{preSum: make([][]int, m+1)}
	for i := range numMatrix.preSum {
		numMatrix.preSum[i] = make([]int, n+1)
	}
	for i := range numMatrix.preSum {
		for j := range numMatrix.preSum[i] {
			if i == 0 || j == 0 {
				numMatrix.preSum[i][j] = 0
			} else {
				numMatrix.preSum[i][j] = numMatrix.preSum[i-1][j] + numMatrix.preSum[i][j-1] - numMatrix.preSum[i-1][j-1] + matrix[i-1][j-1]
			}
		}
	}
	return numMatrix
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row1][col1] + this.preSum[row2+1][col2+1] - this.preSum[row2+1][col1] - this.preSum[row1][col2+1]
}
