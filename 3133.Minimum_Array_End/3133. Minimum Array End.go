package main

func minEnd(n, x int) int64 {
	n--
	for i, j := 0, 0; n>>j > 0; i++ { // 思路：将 n 的二进制位从低到高填入 x 的零位
		if x>>i&1 == 0 { // x 的第 i 个比特位为 0
			x |= n >> j & 1 << i // 在 n 的第 j 个比特位填入当前空位
			j++
		}
	}
	return int64(x)
}
