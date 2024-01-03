package main

// 本质：判断有向图是否存在环
// 见第210题
// DFS
//func canFinish(numCourses int, prerequisites [][]int) bool {
//	var edges, visited, valid = make([][]int, numCourses), make([]int, numCourses), true
//	var result []int
//	var dfs func(int)
//	for _, prerequisite := range prerequisites {
//		edges[prerequisite[1]] = append(edges[prerequisite[1]], prerequisite[0])
//	}
//	dfs = func(u int) {
//		visited[u] = 1
//		for _, v := range edges[u] {
//			switch visited[v] {
//			case 0:
//				dfs(v)
//				if !valid {
//					return
//				}
//			case 1:
//				valid = false
//				return
//			}
//		}
//		visited[u], result = 2, append(result, u)
//	}
//	for i := 0; i < numCourses && valid; i++ {
//		if visited[i] == 0 {
//			dfs(i)
//		}
//	}
//	return valid
//}

// BFS
func canFinish(numCourses int, prerequisites [][]int) bool {
	var edges = make([][]int, numCourses)
	var indeg = make([]int, numCourses) // 维护每个节点的入度
	var seq, q []int
	for _, prerequisite := range prerequisites {
		edges[prerequisite[1]] = append(edges[prerequisite[1]], prerequisite[0])
		indeg[prerequisite[0]]++
	}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		var u = q[0]
		q = q[1:]
		seq = append(seq, u)
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return len(seq) == numCourses // 如果无法构造一个包含所有课程且满足学习顺序的序列
}
