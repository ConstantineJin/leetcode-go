package main

func smallestMissingValueSubtree(parents []int, nums []int) []int {
	n := len(parents)
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		children[parents[i]] = append(children[parents[i]], i)
	}
	geneSet, visited := make(map[int]bool), make([]bool, n)
	var dfs func(int)
	dfs = func(node int) {
		if visited[node] {
			return
		}
		visited[node], geneSet[nums[node]] = true, true
		for _, child := range children[node] {
			dfs(child)
		}
	}
	ans, node, iNode := make([]int, n), -1, 1
	for i := 0; i < n; i++ {
		ans[i] = 1
		if nums[i] == 1 {
			node = i
		}
	}
	for node != -1 {
		dfs(node)
		for geneSet[iNode] {
			iNode++
		}
		ans[node], node = iNode, parents[node]
	}
	return ans
}
