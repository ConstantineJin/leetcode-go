package main

import "math/bits"

func findIntegers(n int) int {
	m := bits.Len(uint(n))
	mem := make([][2]int, m)
	for i := range mem {
		mem[i] = [2]int{-1, -1}
	}
	var dfs func(i, pre int, isLimit bool) int // dfs(i, pre, isLimit) 表示当前处理从左往右第 i 位，前一位为 pre，isLimit 为 true 表示当前位受到 n 的约束，至多填入 n >> i & 1
	dfs = func(i, pre int, isLimit bool) (res int) {
		if i < 0 {
			return 1
		}
		if !isLimit { // 只在不受限制时进行记忆化。因为对于一对 (i, pre) 在整个递归过程中至多受到一次约束，无需记忆化受约束的搜索结果
			if v := mem[i][pre]; v != -1 {
				return v
			}
			defer func() { mem[i][pre] = res }()
		}
		up := 1 // 当前位的取值上限
		if isLimit {
			up = n >> i & 1 // n 中这一位的值
		}
		res = dfs(i-1, 0, isLimit && up == 0) // 填 0
		if pre == 0 && up == 1 {              // 可以填 1
			res += dfs(i-1, 1, isLimit)
		}
		return
	}
	return dfs(m-1, 0, true) // 从高位到低位
}
