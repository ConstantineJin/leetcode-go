package main

type UFS struct{ parent, size []int }

func newUFS(n int) *UFS {
	parent, size := make([]int, n), make([]int, n)
	for i := range parent {
		parent[i], size[i] = i, 1
	}
	return &UFS{parent: parent, size: size}
}

func (u *UFS) find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}

func (u *UFS) isSameSet(x, y int) bool {
	return u.find(x) == u.find(y)
}

func (u *UFS) union(x, y int) {
	var fx, fy = u.find(x), u.find(y)
	if fx == fy {
		return
	}
	if u.size[fx] < u.size[fy] {
		fx, fy = fy, fx
	}
	u.size[fx] += u.size[fy]
	u.parent[fy] = fx
}

func equationsPossible(equations []string) bool {
	var equals, unequals [][2]int
	for _, equation := range equations {
		a, b := int(equation[0]-'a'), int(equation[3]-'a')
		if equation[1] == '=' {
			equals = append(equals, [2]int{a, b})
		} else {
			if a == b {
				return false
			}
			unequals = append(unequals, [2]int{a, b})
		}
	}
	ufs := newUFS(26)
	for _, equal := range equals {
		ufs.union(equal[0], equal[1])
	}
	for _, unequal := range unequals {
		if ufs.isSameSet(unequal[0], unequal[1]) {
			return false
		}
	}
	return true
}
