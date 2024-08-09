package main

func numMagicSquaresInside(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 5 {
				a, b, c, d, e, f, g, h := grid[i-1][j-1], grid[i-1][j], grid[i-1][j+1], grid[i][j-1], grid[i][j+1], grid[i+1][j-1], grid[i+1][j], grid[i+1][j+1]
				if a%5 != 0 && b%5 != 0 && c%5 != 0 && d%5 != 0 && a+h == 10 && c+f == 10 && b+g == 10 && d+e == 10 && a+b+c == 15 && f+g+h == 15 && a+d+f == 15 && c+e+h == 15 {
					ans++
				}
			}
		}
	}
	return
}
