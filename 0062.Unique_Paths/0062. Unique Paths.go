package main

import "math/big"

func uniquePaths(m int, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64()) // 总共需要向下移动m-1次，向右移动n-1次，因此方案数是从m+n-2次选择m-1次的组合数
}
