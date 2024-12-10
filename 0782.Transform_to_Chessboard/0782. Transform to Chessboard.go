package main

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func movesToChessboard(board [][]int) int {
	n := len(board)
	firstRow, firstCol := board[0], make([]int, n)
	var rowCnt, colCnt [2]int
	for i, row := range board {
		rowCnt[firstRow[i]]++
		firstCol[i] = row[0]
		colCnt[firstCol[i]]++
	}
	if abs(rowCnt[0]-rowCnt[1]) > 1 || abs(colCnt[0]-colCnt[1]) > 1 { // 第一行和第一列的 0 和 1 的个数相差不能超过 1
		return -1
	}
	for _, row := range board { // 每一行和第一行比较，要么完全相同要么完全不同，否则返回 -1
		same := row[0] == firstRow[0]
		for i, v := range row {
			if (v == firstRow[i]) != same {
				return -1
			}
		}
	}
	return minSwap(firstRow, rowCnt) + minSwap(firstCol, colCnt)
}

func minSwap(s []int, cnt [2]int) int {
	var x0, diff int // 如果 n 是偶数，x0 是 0
	if cnt[1] > cnt[0] {
		x0 = 1
	}
	for i, x := range s {
		diff += x ^ i%2 ^ x0
	}
	n := len(s)
	if n%2 > 0 {
		return diff / 2
	}
	return min(diff, n-diff) / 2
}
