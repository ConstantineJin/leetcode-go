package main

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// inCircle 判断点 (x, y) 是否在 以 (ox, oy) 为圆心的圆中（含边界）
func inCircle(ox, oy, r, x, y int) bool {
	return (ox-x)*(ox-x)+(oy-y)*(oy-y) <= r*r
}

func canReachCorner(X, Y int, circles [][]int) bool {
	vis := make([]bool, len(circles))
	var dfs func(i int) bool
	dfs = func(i int) bool {
		x1, y1, r1 := circles[i][0], circles[i][1], circles[i][2]
		if y1 <= Y && abs(x1-X) <= r1 || x1 <= X && y1 <= r1 || x1 > X && inCircle(x1, y1, r1, X, 0) { // 圆 i 是否与矩形右/下边界相交
			return true
		}
		vis[i] = true
		for j, circle := range circles {
			x2, y2, r2 := circle[0], circle[1], circle[2]
			// 取两圆心连线上的点 A 满足 |O1A| / |O1O2| = r1 / (r1+r2)。如果两圆相交或相切则点 A 必在交集中。
			// 其横坐标为 (x1·r2+x2·r1)/(r1+r2)，纵坐标为 (y1·r2+y2·r1)/(r1+r2)。若点 A 严格在矩形内（不含边界），则两圆相连。
			if !vis[j] && (x1-x2)*(x1-x2)+(y1-y2)*(y1-y2) <= (r1+r2)*(r1+r2) && x1*r2+x2*r1 < (r1+r2)*X && y1*r2+y2*r1 < (r1+r2)*Y && dfs(j) {
				return true
			}
		}
		return false
	}
	for i, circle := range circles {
		x, y, r := circle[0], circle[1], circle[2]
		if inCircle(x, y, r, 0, 0) || inCircle(x, y, r, X, Y) || !vis[i] && (x <= X && abs(y-Y) <= r || y <= Y && x <= r || y > Y && inCircle(x, y, r, 0, Y)) && dfs(i) { // 从与矩形左/上边界相交的圆开始 DFS
			return false
		}
	}
	return true
}
