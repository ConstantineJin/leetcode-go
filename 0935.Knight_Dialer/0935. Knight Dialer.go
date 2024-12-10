package main

const mod int = 1e9 + 7

// var next = [10][]int{{4, 6}, {6, 8}, {7, 9}, {4, 8}, {0, 3, 9}, {}, {0, 1, 7}, {2, 6}, {1, 3}, {2, 4}}
var f = [5000][4]int{{1, 1, 1, 1}} // 从状态转移角度考虑，10 个键位可以分为四类，分别为 {0}, {1,3,7,9}, {2,8}, {4,6}

func init() {
	for i := 1; i < len(f); i++ {
		f[i][0] = f[i-1][3] * 2 % mod
		f[i][1] = (f[i-1][2] + f[i-1][3]) % mod
		f[i][2] = f[i-1][1] * 2 % mod
		f[i][3] = (f[i-1][1]*2 + f[i-1][0]) % mod
	}
}

func knightDialer(n int) int {
	if n == 1 {
		return 10
	}
	return (f[n-1][0] + f[n-1][1]*4 + f[n-1][2]*2 + f[n-1][3]*2) % mod
}
