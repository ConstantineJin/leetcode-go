package main

func minFlips(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m/2; i++ {
		for j := 0; j < n/2; j++ {
			topLeft, topRight, bottomLeft, bottomRight := grid[i][j], grid[i][n-j-1], grid[m-i-1][j], grid[m-i-1][n-j-1]
			s := topLeft + topRight + bottomLeft + bottomRight
			ans += min(s, 4-s)
		}
	}
	rowMid, colMid := m/2, n/2
	var a, b int // a 表示对称位置值不相等的数量，b 表示对称位置值均为 1 的数量
	if m%2 == 1 && n%2 == 1 {
		ans += grid[rowMid][colMid] // 正中间必须是 0
		for i := 0; i < rowMid; i++ {
			if grid[i][colMid] != grid[m-i-1][colMid] {
				a++
				ans++
			} else if grid[i][colMid] == 1 {
				b++
			}
		}
		for j := 0; j < colMid; j++ {
			if grid[rowMid][j] != grid[rowMid][n-j-1] {
				a++
				ans++
			} else if grid[rowMid][j] == 1 {
				b++
			}
		}
		if b%2 == 1 && a == 0 { // 要满足 1 的数量能被 4 整除，必须将一组对称的 0 都变成 1
			ans += 2
		}
	} else if m%2 == 1 {
		for j := 0; j < colMid; j++ {
			if grid[rowMid][j] != grid[rowMid][n-j-1] {
				a++
				ans++
			} else if grid[rowMid][j] == 1 {
				b++
			}
		}
		if b%2 == 1 && a == 0 {
			ans += 2
		}
	} else if n%2 == 1 {
		for i := 0; i < rowMid; i++ {
			if grid[i][colMid] != grid[m-i-1][colMid] {
				a++
				ans++
			} else if grid[i][colMid] == 1 {
				b++
			}
		}
		if b%2 == 1 && a == 0 {
			ans += 2
		}
	}
	return
}
