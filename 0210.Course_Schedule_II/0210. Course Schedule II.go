package main

// 拓扑排序
// DFS
//func findOrder(numCourses int, prerequisites [][]int) []int {
//	var seq []int                         // 学习顺序
//	var visited = make([]int, numCourses) // 每个节点的访问状态，0表示未被搜索，1表示正在搜索，2表示搜索完成
//	var edges = make([][]int, numCourses) // 记录每个节点的邻接节点
//	var valid = true                      // 是否能找到符合要求的学习顺序
//	var dfs func(int)
//	for _, prerequisite := range prerequisites { // 遍历先修关系
//		edges[prerequisite[1]] = append(edges[prerequisite[1]], prerequisite[0]) // 初始化邻接表
//	}
//	dfs = func(i int) {
//		visited[i] = 1               // 将当前节点标记为正在搜索中
//		for _, v := range edges[i] { // 遍历当前节点的所有相邻节点
//			switch visited[v] {
//			case 0: // 当前节点未被访问过
//				dfs(v)
//				if !valid {
//					return
//				}
//			case 1: // 当前节点也在搜索中，说明有环
//				valid = false
//				return
//			}
//		}
//		visited[i], seq = 2, append(seq, i) // 当前节点的所有邻接节点都已完成搜索，则当且节点也标记为完成搜索，将其入栈
//	}
//	for i := 0; i < numCourses && valid; i++ { // 对每一个未访问的节点进行dfs
//		if visited[i] == 0 {
//			dfs(i)
//		}
//	}
//	if !valid { // 如果找不到符合要求的学习顺序
//		return []int{} // 依题意返回空数组
//	}
//	for i := 0; i < numCourses/2; i++ {
//		seq[i], seq[numCourses-i-1] = seq[numCourses-i-1], seq[i] // 反转seq
//	}
//	return seq
//}

// BFS
func findOrder(numCourses int, prerequisites [][]int) []int {
	var edges = make([][]int, numCourses)
	var indeg = make([]int, numCourses)          // 维护每个节点的入度
	var seq, q []int                             // q维护当前入度为0的节点队列
	for _, prerequisite := range prerequisites { // 根据先修关系初始化邻接表和入度数组
		edges[prerequisite[1]] = append(edges[prerequisite[1]], prerequisite[0])
		indeg[prerequisite[0]]++
	}
	for i := 0; i < numCourses; i++ { // 初始化q
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 { // 只要队列不空
		var u = q[0] // 取出队首
		q = q[1:]
		seq = append(seq, u)         // 添加到顺序中
		for _, v := range edges[u] { // 遍历原队首节点的邻接节点
			indeg[v]--         // 将其入度-1
			if indeg[v] == 0 { // 如果有节点的入度变为0
				q = append(q, v) // 将其添加到队列中
			}
		}
	}
	if len(seq) != numCourses { // 如果无法构造一个包含所有课程且满足学习顺序的序列
		return []int{} // 依题意返回空数组
	}
	return seq
}
