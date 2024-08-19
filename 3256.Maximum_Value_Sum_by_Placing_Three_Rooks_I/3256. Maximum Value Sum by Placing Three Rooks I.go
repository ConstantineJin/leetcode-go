package main

import "math"

type pair struct{ value, j int }

func maximumValueSum(board [][]int) int64 {
	m := len(board)
	suf := make([][3]pair, m) // 前后缀分解，先预处理后缀
	p := [3]pair{}            // 维护前三大
	for i := range p {
		p[i].value = math.MinInt
	}
	update := func(row []int) {
		for j, v := range row {
			if v > p[0].value {
				if p[0].j != j { // 如果相等，仅更新最大
					if p[1].j != j { // 如果相等，仅更新最大和次大
						p[2] = p[1]
					}
					p[1] = p[0]
				}
				p[0] = pair{v, j}
			} else if v > p[1].value && j != p[0].j {
				if p[1].j != j { // 如果相等，仅更新最大
					p[2] = p[1]
				}
				p[1] = pair{v, j}
			} else if v > p[2].value && j != p[0].j && j != p[1].j {
				p[2] = pair{v, j}
			}
		}
	}
	for i := m - 1; i > 1; i-- {
		update(board[i])
		suf[i] = p
	}
	ans := math.MinInt
	for i := range p {
		p[i].value = math.MinInt // 重置，计算前缀
	}
	for i, row := range board[:m-2] {
		update(row)
		for j, x := range board[i+1] { // 枚举中间的车
			for _, p := range p { // 枚举第一辆车
				for _, q := range suf[i+2] { // 枚举第三辆车
					if p.j != j && q.j != j && p.j != q.j { // 都不在同一列
						ans = max(ans, p.value+x+q.value)
						break
					}
				}
			}
		}
	}
	return int64(ans)
}
