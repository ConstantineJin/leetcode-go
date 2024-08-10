package main

type UnionFindSet struct {
	parent, size []int
	count        int
}

func newUnionFindSet(n int) *UnionFindSet {
	parent, size := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		parent[i], size[i] = i, 1
	}
	return &UnionFindSet{parent, size, n}
}

func (u *UnionFindSet) Find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.Find(u.parent[x])
	}
	return u.parent[x]
}

func (u *UnionFindSet) Union(x, y int) {
	fx, fy := u.Find(x), u.Find(y)
	if fx == fy {
		return
	}
	if u.size[fx] < u.size[fy] {
		fx, fy = fy, fx
	}
	u.size[fx] += u.size[fy]
	u.parent[fy] = fx
	u.count--
}

func regionsBySlashes(grid []string) int {
	n := len(grid)
	ufs := newUnionFindSet(n * n * 4) // 将每个方格的对角线划分出的四个小三角形视为基本单元，按上右下左的顺序从 0～3 标记，判断其连通性
	for i, row := range grid {
		for j, v := range row {
			idx := i*n + j
			if i < n-1 { // 如果不是最后一行，就把当前单元格的底部三角形和下方单元格的顶部三角形合并
				bottom := idx + n
				ufs.Union(idx*4+2, bottom*4)
			}
			if j < n-1 { // 如果不是最后一列，就把当前单元格的右侧三角形和右侧单元格的左侧三角形合并
				right := idx + 1
				ufs.Union(idx*4+1, right*4+3)
			}
			switch v {
			case '/':
				ufs.Union(idx*4, idx*4+3)
				ufs.Union(idx*4+1, idx*4+2)
			case '\\':
				ufs.Union(idx*4, idx*4+1)
				ufs.Union(idx*4+2, idx*4+3)
			default:
				ufs.Union(idx*4, idx*4+1)
				ufs.Union(idx*4+1, idx*4+2)
				ufs.Union(idx*4+2, idx*4+3)
			}
		}
	}
	return ufs.count
}
