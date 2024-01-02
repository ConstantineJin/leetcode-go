package main

// 多源BFS
func orangesRotting(grid [][]int) (ans int) {
	var (
		m, n = len(grid), len(grid[0]) // 网格的长宽
		cur  [][2]int                  // 当前遍历的腐烂橘子的坐标
		good int                       // 好橘子的数量
	)
	for i, row := range grid { // 初始化good和cur
		for j, v := range row {
			switch v {
			case 1:
				good++
			case 2:
				cur = append(cur, [2]int{i, j})
			}
		}
	}
	for len(cur) > 0 { // 只要还有刚被腐烂的橘子
		var nxt = make([][2]int, 0) // 下一分钟会被腐烂的橘子坐标
		for _, pos := range cur {   // 遍历刚被腐烂的橘子坐标
			if pos[0] > 0 && grid[pos[0]-1][pos[1]] == 1 { // 遍历每个橘子的上下左右网格，如果是好橘子
				grid[pos[0]-1][pos[1]] = 2                    // 腐烂该橘子
				nxt = append(nxt, [2]int{pos[0] - 1, pos[1]}) // 将其添加到下一腐烂的橘子队列
				good--                                        // 好橘子数量-1
			}
			if pos[1] > 0 && grid[pos[0]][pos[1]-1] == 1 {
				grid[pos[0]][pos[1]-1] = 2
				nxt = append(nxt, [2]int{pos[0], pos[1] - 1})
				good--
			}
			if pos[0] < m-1 && grid[pos[0]+1][pos[1]] == 1 {
				grid[pos[0]+1][pos[1]] = 2
				nxt = append(nxt, [2]int{pos[0] + 1, pos[1]})
				good--
			}
			if pos[1] < n-1 && grid[pos[0]][pos[1]+1] == 1 {
				grid[pos[0]][pos[1]+1] = 2
				nxt = append(nxt, [2]int{pos[0], pos[1] + 1})
				good--
			}
		}
		if len(nxt) > 0 { // 有腐蚀新的橘子
			ans++ // 分钟数+1
		}
		cur = nxt // 进入下一分钟
	}
	if good > 0 { // 仍有好橘子无法被腐蚀
		return -1 // 返回-1
	}
	return
}
