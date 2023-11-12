package main

// 方法1：并查集
// 将N对情侣看做图中的N个节点；对于每对相邻的位置，如果是第i对与第j对坐在了一起，则在i号节点与j号节点之间连接一条边，代表需要交换这两对情侣的位置
// 如果图中形成了一个大小为k的环，则沿着环的方向，进行了k−1次交换后，这k对情侣就都能够彼此牵手了
// 故我们只需要利用并查集求出图中的每个连通分量；对于每个连通分量而言，其大小减1就是需要交换的次数

//// UFS 并查集
//type UFS struct {
//	parent, size []int
//	setCount     int // 当前集合数量，即连通分量数量
//}
//
//// newUFS 构造新的并查集
//func newUFS(n int) *UFS {
//	parent, size := make([]int, n), make([]int, n)
//	for i := range parent {
//		parent[i], size[i] = i, 1
//	}
//	return &UFS{parent: parent, size: size, setCount: n}
//}
//
//// find 寻找根
//func (u *UFS) find(x int) int {
//	if u.parent[x] != x {
//		u.parent[x] = u.find(u.parent[x])
//	}
//	return u.parent[x]
//}
//
//// union 合并两个集合
//func (u *UFS) union(x, y int) {
//	fx, fy := u.find(x), u.find(y)
//	if fx == fy { // 根相同，说明已在同一个集合中
//		return
//	}
//	if u.size[fx] < u.size[fy] { // 否则进行合并。适当优化
//		fx, fy = fy, fx
//	}
//	u.size[fx] += u.size[fy]
//	u.parent[fy] = fx
//	u.setCount--
//}
//
//func minSwapsCouples(row []int) int {
//	n := len(row)
//	ufs := newUFS(n / 2)
//	for i := 0; i < n; i += 2 {
//		ufs.union(row[i]/2, row[i+1]/2)
//	}
//	return n/2 - ufs.setCount
//}

// 方法2：广度优先搜索
func minSwapsCouples(row []int) (res int) {
	n := len(row)                                       // 总人数，即情侣对数*2
	graph, vis := make([][]int, n/2), make([]bool, n/2) // graph记录连通，vis记录是否已访问
	for i := 0; i < n; i += 2 {                         // 顺序遍历rows，每次查看两个人
		l, r := row[i]/2, row[i+1]/2
		if l != r {
			graph[l], graph[r] = append(graph[l], r), append(graph[r], l) // 相邻座不是一对情侣，则连通两人的所属节点
		}
	}
	for i, v := range vis {
		if !v { // 如果未访问
			vis[i] = true
			cnt, q := 0, []int{i}
			for len(q) > 0 {
				cnt++
				v := q[0]
				q = q[1:]
				for _, vv := range graph[v] {
					if !vis[vv] {
						vis[vv], q = true, append(q, vv)
					}
				}
			}
			res += cnt - 1
		}
	}
	return
}
