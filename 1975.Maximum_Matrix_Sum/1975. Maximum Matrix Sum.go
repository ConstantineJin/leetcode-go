package main

import "math"

func maxMatrixSum(matrix [][]int) int64 {
	var negativeCnt, ans int
	mn := math.MaxInt
	for _, row := range matrix {
		for _, v := range row {
			if v < 0 {
				negativeCnt++
				mn = min(mn, -v)
				ans -= v
			} else {
				mn = min(mn, v)
				ans += v
			}
		}
	}
	if negativeCnt&1 == 1 {
		ans -= mn << 1
	}
	return int64(ans)
}
