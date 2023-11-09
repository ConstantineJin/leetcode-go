package main

type pair struct{ x, y int }

var dirs = []pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func maximumMinutes(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	bfs := func(q []pair) (int, int, int) { // bfs返回三个数，分别是人/火到达安全屋、安全屋左侧和安全屋上方格子的最短时间
		time := make([][]int, m)
		for i := range time {
			time[i] = make([]int, n)
			for j := range time[i] {
				time[i][j] = -1 // -1表示未访问过该节点
			}
		}
		for _, p := range q {
			time[p.x][p.y] = 0 // 起始点的抵达时间是0
		}
		for t := 1; len(q) > 0; t++ { // bfs求得抵达每一个点的时间
			tmp := q
			q = nil
			for _, p := range tmp {
				for _, d := range dirs {
					if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == 0 && time[x][y] < 0 {
						time[x][y] = t
						q = append(q, pair{x, y})
					}
				}
			}
		}
		return time[m-1][n-1], time[m-1][n-2], time[m-2][n-1]
	}

	manToHouse, manToLeft, manToUp := bfs([]pair{{0, 0}}) // 人到达安全屋、安全屋左侧的格子和安全屋上方的格子的时间
	if manToHouse < 0 {                                   // 人无法到达安全屋
		return -1
	}

	var firePos []pair // 起始着火点的位置
	for i, row := range grid {
		for j, x := range row {
			if x == 1 {
				firePos = append(firePos, pair{i, j})
			}
		}
	}

	fireToHouse, fireToLeft, fireToUp := bfs(firePos) // 火到达安全屋、安全屋左侧的格子和安全屋上方的格子的时间
	if fireToHouse < 0 {                              // 火无法抵达安全屋
		return 1e9
	}

	delta := fireToHouse - manToHouse // 火与人到达安全屋的时间差
	if delta < 0 {                    // 火比人先到安全屋
		return -1
	}
	if manToLeft != -1 && manToLeft+delta < fireToLeft || manToUp != -1 && manToUp+delta < fireToUp { // 人比火先到达安全屋左侧或上方的格子
		return delta
	}
	return delta - 1 // 如果多等最后一分钟安全屋已被火包围，无法进入，需减去（人与火可以同时抵达安全屋）
}
