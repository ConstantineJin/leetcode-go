package main

const mod int = 1e9 + 7

func checkRecord(n int) (ans int) {
	f := make([][2][3]int, n+1) // 三个维度分别表示当前天数、累计缺席天数、连续迟到天数
	f[0][0][0] = 1
	for i := 1; i <= n; i++ {
		// 到场
		for j := 0; j < 2; j++ {
			for k := 0; k < 3; k++ {
				f[i][j][0] = (f[i][j][0] + f[i-1][j][k]) % mod
			}
		}
		// 缺席
		for k := 0; k < 3; k++ {
			f[i][1][0] = (f[i][1][0] + f[i-1][0][k]) % mod
		}
		// 迟到
		for j := 0; j < 2; j++ {
			for k := 1; k < 3; k++ {
				f[i][j][k] = (f[i][j][k] + f[i-1][j][k-1]) % mod
			}
		}
	}
	for j := 0; j < 2; j++ {
		for k := 0; k < 3; k++ {
			ans = (ans + f[n][j][k]) % mod
		}
	}
	return ans
}
