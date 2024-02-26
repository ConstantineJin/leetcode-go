package main

type UFS struct{ parent, size []int }

func newUFS(n int) *UFS {
	var parent, size = make([]int, n), make([]int, n)
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

func (u *UFS) union(x, y int) {
	var fx, fy = u.find(x), u.find(y)
	if fx != fy {
		if u.size[fx] < u.size[fy] {
			fx, fy = fy, fx
		}
		u.size[fx] += u.size[fy]
		u.parent[fy] = fx
	}
}

func canTraverseAllPairs(nums []int) bool {
	var n = len(nums)
	var mp = make(map[int][]int) // key为质因数，value为nums中含有该因数的元素下标
	for i, num := range nums {
		for x := 2; x*x <= num; x++ {
			if num%x == 0 {
				mp[x] = append(mp[x], i)
				for num /= x; num%x == 0; num /= x {
				}
			}
		}
		if num > 1 {
			mp[num] = append(mp[num], i)
		}
	}
	var ufs = newUFS(n)
	for _, v := range mp {
		for i := len(v) - 1; i > 0; i-- {
			ufs.union(v[0], v[i])
		}
	}
	return ufs.size[ufs.find(0)] == n
}
