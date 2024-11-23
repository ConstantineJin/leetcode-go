package main

func rotateTheBox(box [][]byte) [][]byte {
	m, n := len(box), len(box[0])
	ans := make([][]byte, n)
	for i := range ans {
		ans[i] = make([]byte, m)
		for j := range ans[i] {
			ans[i][j] = '.'
		}
	}
	for i, row := range box {
		for j := 0; j < n; j++ {
			var cnt int
			for ; j < n && row[j] != '*'; j++ {
				if row[j] == '#' {
					cnt++
				}
			}
			if j < n {
				ans[j][m-i-1] = '*'
			}
			for k := j - 1; cnt > 0; k-- {
				ans[k][m-i-1] = '#'
				cnt--
			}
		}
	}
	return ans
}
