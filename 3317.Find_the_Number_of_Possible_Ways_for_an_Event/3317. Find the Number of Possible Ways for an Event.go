package main

const mod int = 1e9 + 7
const mx = 1001

var s [mx][mx]int // s[i][j] 表示把 i 个人划分成 j 个非空集合的方案数

func init() {
	s[0][0] = 1
	for i := 1; i < mx; i++ {
		for j := 1; j <= i; j++ {
			s[i][j] = (s[i-1][j-1] + j*s[i-1][j]) % mod // 加法原理，第 i 个人独立成一个集合（转移到 s[i-1][j-1])或者分配到前 j 个集合中的一个（转移到 s[i-1][j]，共 j 个）
		}
	}
}

func numberOfWays(n, x, y int) (ans int) {
	perm, powY := 1, 1
	for i := 1; i <= min(n, x); i++ { // 枚举有 i 个有表演者的节目
		perm = perm * (x + 1 - i) % mod           // 从 x 个节目中有序地选择 i 个节目的方案数
		powY = powY * y % mod                     // 有 i 个节目需要打分，每个节目有 y 种取值，共有 y^i 种
		ans = (ans + perm*s[n][i]%mod*powY) % mod // 乘法原理，把 n 个人划分成 i 个非空集合的方案数
	}
	return
}
