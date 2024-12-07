package main

var dirs = [8][2]int{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}

func knightProbability(n, k, row, column int) float64 {
	f := make([][][]float64, k+1)
	for i := range f {
		f[i] = make([][]float64, n+4)
		for j := range f[i] {
			f[i][j] = make([]float64, n+4)
		}
	}
	for i := 2; i < n+2; i++ {
		for j := 2; j < n+2; j++ {
			f[0][i][j] = 1
		}
	}
	for steps := 1; steps <= k; steps++ {
		for i := 2; i < n+2; i++ {
			for j := 2; j < n+2; j++ {
				for _, dir := range dirs {
					f[steps][i][j] += f[steps-1][i+dir[0]][j+dir[1]]
				}
				f[steps][i][j] /= 8
			}
		}
	}
	return f[k][row+2][column+2]
}
