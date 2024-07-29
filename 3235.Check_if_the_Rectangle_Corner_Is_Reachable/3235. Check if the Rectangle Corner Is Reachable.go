package main

func canReachCorner(X, Y int, circles [][]int) bool {
	n := len(circles)
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y int) {
		rootX, rootY := find(x), find(y)
		if rootX != rootY {
			parent[rootY] = rootX
		}
	}

	areCirclesConnected := func(c1, c2 []int) bool {
		dx, dy, r := c1[0]-c2[0], c1[1]-c2[1], c1[2]+c2[2]
		return dx*dx+dy*dy <= r*r
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if areCirclesConnected(circles[i], circles[j]) {
				union(i, j)
			}
		}
	}

	touchTop, touchBottom, touchLeft, touchRight := make(map[int]bool), make(map[int]bool), make(map[int]bool), make(map[int]bool)
	for i, circle := range circles {
		x, y, r := circle[0], circle[1], circle[2]
		root := find(i)
		if y+r >= Y {
			touchTop[root] = true
		}
		if y-r <= 0 {
			touchBottom[root] = true
		}
		if x-r <= 0 {
			touchLeft[root] = true
		}
		if x+r >= X {
			touchRight[root] = true
		}
	}

	for i := 0; i < n; i++ {
		root := find(i)
		if touchTop[root] && touchBottom[root] || touchLeft[root] && touchRight[root] || touchLeft[root] && touchBottom[root] || touchTop[root] && touchRight[root] {
			return false
		}
	}

	return true
}
