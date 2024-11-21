package main

func countUnguarded(m, n int, guards, walls [][]int) (ans int) {
	grids := make([][]int, m)
	for i := range grids {
		grids[i] = make([]int, n)
	}
	for _, guard := range guards {
		grids[guard[0]][guard[1]] = 1
	}
	for _, wall := range walls {
		grids[wall[0]][wall[1]] = 2
	}
	for _, row := range grids {
		for j := 0; j < n; j++ {
			if row[j] == 2 {
				continue
			}
			st, has1 := j, false
			for ; j < n && row[j] != 2; j++ {
				if row[j] == 1 {
					has1 = true
				}
			}
			if has1 {
				for ; st < j; st++ {
					if row[st] == 0 {
						row[st] = -1
					}
				}
			}
		}
	}
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if grids[i][j] == 2 {
				continue
			}
			st, has1 := i, false
			for ; i < m && grids[i][j] != 2; i++ {
				if grids[i][j] == 1 {
					has1 = true
				}
			}
			if has1 {
				for ; st < i; st++ {
					if grids[st][j] == 0 {
						grids[st][j] = -1
					}
				}
			}
		}
	}
	for _, row := range grids {
		for _, v := range row {
			if v == 0 {
				ans++
			}
		}
	}
	return
}
