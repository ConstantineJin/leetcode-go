package main

type pair struct{ v, d int }

func secondMinimum(n int, edges [][]int, time int, change int) int {
	g := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0]-1, edge[1]-1
		g[x], g[y] = append(g[x], y), append(g[y], x)
	}
	next := func(d int) int { // 输入当前时间 d，返回到达下一个节点的时间
		if times := d / change; times%2 == 1 { // 当前为红灯
			return (times+1)*change + time // 等绿灯再出发
		}
		return d + time // 绿灯，直接出发
	}
	dis := make([][2]int, n) // 最短路和次短路
	dis[0][1] = 1e9
	for i := 1; i < n; i++ {
		dis[i] = [2]int{1e9, 1e9}
	}
	q := []pair{{}}
	for len(q) > 0 { // BFS
		p := q[0]
		q = q[1:]
		for _, w := range g[p.v] {
			d := next(p.d)     // 到达节点 w 的时间
			if d < dis[w][0] { // 更新最短路
				dis[w][0] = d
				q = append(q, pair{w, d})
			} else if dis[w][0] < d && d < dis[w][1] { // 更新次短路
				dis[w][1] = d
				q = append(q, pair{w, d})
			}
		}
	}
	return dis[n-1][1]
}
