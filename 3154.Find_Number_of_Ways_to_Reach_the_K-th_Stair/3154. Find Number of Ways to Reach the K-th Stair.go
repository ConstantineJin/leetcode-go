package main

import "math/bits"

// 方法一：记忆化搜索
//func waysToReachStair(k int) int {
//	type args struct {
//		i, jump   int
//		downAvail bool
//	}
//	mem := make(map[args]int)
//	var dfs func(i, jump int, downAvail bool) int
//	dfs = func(i, jump int, downAvail bool) int {
//		if i > k+1 {
//			return 0
//		}
//		p := args{i, jump, downAvail}
//		if v, ok := mem[p]; ok {
//			return v
//		}
//		res := dfs(i+1<<jump, jump+1, true)
//		if downAvail && i > 0 {
//			res += dfs(i-1, jump, false)
//		}
//		if i == k {
//			res++
//		}
//		mem[p] = res
//		return res
//	}
//	return dfs(1, 0, true)
//}

// 方法二：组合数学
const mx = 31

var c [mx][mx]int

func init() {
	for i := 0; i < mx; i++ {
		c[i][0], c[i][i] = 1, 1
		for j := 1; j < i; j++ {
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
	}
}

func waysToReachStair(k int) (ans int) {
	for j := bits.Len(uint(max(k-1, 0))); 1<<j-k <= j+1; j++ {
		ans += c[j+1][1<<j-k]
	}
	return
}
