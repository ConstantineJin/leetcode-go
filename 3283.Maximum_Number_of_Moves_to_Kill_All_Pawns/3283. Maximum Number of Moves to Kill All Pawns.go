package main

import (
	"math"
	"math/bits"
)

type pair struct{ x, y int }

var directions = [8][2]int{{1, 2}, {1, -2}, {-1, 2}, {-1, -2}, {2, 1}, {2, -1}, {-2, 1}, {-2, -1}}

func maxMoves(kx, ky int, positions [][]int) int {
	n := len(positions)                 // 初始状态下棋盘上兵的个数
	distances := make([][50][50]int, n) // distances[i][x][y] 表示马从 (x,y) 移动到第 i 个兵的位置所需要的步数
	for i, position := range positions {
		distance := &distances[i]
		for j := range distance {
			for k := range distances[j] {
				distance[j][k] = -1
			}
		}
		px, py := position[0], position[1]
		distance[px][py] = 0
		cur := []pair{{px, py}} // BFS
		for step := 1; len(cur) > 0; step++ {
			var nxt []pair
			for _, p := range cur {
				for _, direction := range directions {
					x, y := p.x+direction[0], p.y+direction[1]
					if 0 <= x && x < 50 && 0 <= y && y < 50 && distance[x][y] < 0 {
						distance[x][y] = step
						nxt = append(nxt, pair{x, y})
					}
				}
			}
			cur = nxt
		}
	}
	positions = append(positions, []int{kx, ky}) // 把马的位置当作第 n+1 个兵的位置
	mem := make([][]int, n+1)
	for i := range mem {
		mem[i] = make([]int, 1<<n)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	u := 1<<n - 1                       // 全集
	var dfs func(i, mask int) (res int) // 状态压缩DP。dfs(i,mask) 表示当前马在第 i 个兵的位置，还留在棋盘上的兵的集合为 mask，Alice 和 Bob的总移动次数的最大值
	dfs = func(i, mask int) (res int) {
		if mask == 0 { // 棋盘上没有兵了，游戏结束
			return
		}
		if v := mem[i][mask]; v != -1 { // 之前计算过
			return v
		}
		x, y := positions[i][0], positions[i][1]
		if bits.OnesCount(uint(u^mask))%2 == 0 { // Alice 操作时，被吃掉的兵的数量为偶数
			for s := uint(mask); s > 0; s &= s - 1 { // 枚举吃掉哪个兵
				j := bits.TrailingZeros(s)
				res = max(res, dfs(j, mask^1<<j)+distances[j][x][y])
			}
		} else { // Bob 操作时，被吃掉的兵的数量为奇数
			res = math.MaxInt
			for s := uint(mask); s > 0; s &= s - 1 { // 枚举吃掉哪个兵
				j := bits.TrailingZeros(s)
				res = min(res, dfs(j, mask^1<<j)+distances[j][x][y])
			}
		}
		mem[i][mask] = res
		return
	}
	return dfs(n, u)
}
