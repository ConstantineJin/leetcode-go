package main

func stoneGameVII(stones []int) int {
	var n = len(stones)
	var s, mem = make([]int, n+1), make([][]int, n)
	for i, stone := range stones {
		s[i+1] = s[i] + stone // 前缀和
		mem[i] = make([]int, n)
	}
	var dfs func(left, right int) int // dfs(i,j)表示剩余石子从stones[i]到 stones[j]，先手得分减去后手得分的最大值
	dfs = func(left, right int) int {
		if left == right {
			return 0
		}
		var p = &mem[left][right]
		if *p == 0 {
			*p = max(s[right+1]-s[left+1]-dfs(left+1, right), s[right]-s[left]-dfs(left, right-1)) // 枚举移除左侧石子和右侧石子
		}
		return *p
	}
	return dfs(0, n-1)
}
