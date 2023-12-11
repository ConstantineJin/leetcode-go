package main

func minimumFuelCost(roads [][]int, seats int) (ans int64) {
	g := make([][]int, len(roads)+1)
	for _, road := range roads {
		g[road[0]], g[road[1]] = append(g[road[0]], road[1]), append(g[road[1]], road[0]) // 记录每个点的邻居
	}
	var dfs func(int, int) int
	dfs = func(x int, fa int) int {
		size := 1
		for _, y := range g[x] {
			if y != fa { // 递归子节点，不递归父节点
				size += dfs(y, x) // 统计子树大小
			}
		}
		if x > 0 { // x不是根节点
			ans += int64((size-1)/seats + 1) // 需要ceil(size/seats)辆车
		}
		return size
	}
	dfs(0, -1)
	return
}
