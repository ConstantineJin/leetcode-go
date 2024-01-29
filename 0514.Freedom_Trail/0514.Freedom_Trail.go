package main

import "math"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findRotateSteps(ring string, key string) int {
	var m, n = len(ring), len(key)
	var pos = make(map[uint8][]int) // 哈希表pos记录每个字符在ring中的下标
	for i := 0; i < len(ring); i++ {
		pos[ring[i]] = append(pos[ring[i]], i)
	}
	var mem = make([][]int, m) // 记忆化搜索
	for i := range mem {
		mem[i] = make([]int, n)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs func(i, j int) int // i代表当前12点方向字符在ring中的下标，j表示当前匹配的key的下标
	dfs = func(i, j int) (res int) {
		if j == n { // 遍历完成
			return
		}
		var p = &mem[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		res = math.MaxInt
		for _, po := range pos[key[j]] { // 动态规划，遍历当前key[j]字符在ring中所有出现的位置po
			res = min(res, dfs(po, j+1)+min(abs(i-po), m-abs(i-po))) // 结果等于顺/逆时针旋转po到12点位置的较小值，再加上dfs(po,j+1)
		}
		return
	}
	return n + dfs(0, 0) // key中每个字符按下也算一次操作
}
