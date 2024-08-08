package main

func spiralMatrixIII(rows, cols, rStart, cStart int) [][]int {
	total := rows * cols
	ans := make([][]int, 0, total)
	top, bottom, left, right := rStart, rStart, cStart, cStart
	for i := 0; i < total; {
		start1, end1 := max(left, 0), min(right, cols-1)
		for col := start1; col <= end1 && top >= 0 && i < total; col++ {
			ans = append(ans, []int{top, col})
			i++
		}
		right++
		start2, end2 := max(top, 0), min(bottom, rows-1)
		for row := start2; row <= end2 && right < cols && i < total; row++ {
			ans = append(ans, []int{row, right})
			i++
		}
		bottom++
		start3, end3 := min(right, cols-1), max(left, 0)
		for col := start3; col >= end3 && bottom < rows && i < total; col-- {
			ans = append(ans, []int{bottom, col})
			i++
		}
		left--
		start4, end4 := min(bottom, rows-1), max(top, 0)
		for row := start4; row >= end4 && left >= 0 && i < total; row-- {
			ans = append(ans, []int{row, left})
			i++
		}
		top--
	}
	return ans
}
