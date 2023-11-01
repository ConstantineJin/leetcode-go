package main

func maximumInvitations(favorite []int) int {
	n := len(favorite)      //总人数
	indeg := make([]int, n) //统计每个节点的入度，即有多少人喜欢第i个员工
	for _, f := range favorite {
		indeg[f]++
	}

	maxDepth := make([]int, n)
	var q []int //入度为0的点
	for i, d := range indeg {
		if d == 0 {
			q = append(q, i)
		}
	}
	//拓扑排序，剪掉图上所有的枝
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		y := favorite[x]
		maxDepth[y] = maxDepth[x] + 1
		if indeg[y]--; indeg[y] == 0 {
			q = append(q, y)
		}
	}

	maxRingSize, sumChainSize := 0, 0
	for i, d := range indeg {
		if d == 0 {
			continue
		}
		// 遍历基环上的点
		indeg[i] = 0  // 将基环上的点的入度标记为 0，避免重复访问
		ringSize := 1 // 基环长度
		for x := favorite[i]; x != i; x = favorite[x] {
			indeg[x] = 0 // 将基环上的点的入度标记为 0，避免重复访问
			ringSize++
		}

		if ringSize == 2 { // 基环长度为 2
			sumChainSize += maxDepth[i] + maxDepth[favorite[i]] + 2 // 累加两条最长链的长度
		} else {
			maxRingSize = max(maxRingSize, ringSize) // 取所有基环长度的最大值
		}
	}
	return max(maxRingSize, sumChainSize)
}
