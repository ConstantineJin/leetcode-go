package main

import "math/bits"

func maxStudents(seats [][]byte) int {
	var m, n = len(seats), len(seats[0])
	var avail = make([]int, m) // 每排用一个二进制数保存可用椅子的下标
	for i, seat := range seats {
		for j, s := range seat {
			if s == '.' {
				avail[i] |= 1 << j
			}
		}
	}
	var mem = make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, 1<<n)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i int, j int) (res int) {
		var p = &mem[i][j]
		if *p != -1 { // 如果已经计算过了
			return *p // 直接返回结果
		}
		defer func() { *p = res }() // 保存计算结果
		if i == 0 {
			if j == 0 {
				return 0
			}
			var lb = j & -j              // 求出最低位的1
			return dfs(i, j&^(lb*3)) + 1 // j&^(lb*3)把lowbit及其左侧的bit置0
		}
		res = dfs(i-1, avail[i-1])           // 第i排空着
		for s := j; s > 0; s = (s - 1) & j { // 枚举j的子集s
			if s&(s>>1) == 0 { // s没有连续的1
				var t = avail[i-1] &^ (s<<1 | s>>1) // 去掉不能坐人的位置
				res = max(res, dfs(i-1, t)+bits.OnesCount(uint(s)))
			}
		}
		return res
	}
	return dfs(m-1, avail[m-1])
}
