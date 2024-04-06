package main

import "math/bits"

type TreeAncestor [][]int

func Constructor(n int, parent []int) TreeAncestor {
	var m = bits.Len(uint(n))
	var pa = make([][]int, n) // pa[x][i]表示节点x的第2^i个祖先节点。若第2^i个祖先节点不存在，则pa[x][i]=−1
	for i, p := range parent {
		pa[i] = make([]int, m)
		pa[i][0] = p
	}
	for i := 0; i < m-1; i++ {
		for x := 0; x < n; x++ {
			if p := pa[x][i]; p != -1 {
				pa[x][i+1] = pa[p][i]
			} else {
				pa[x][i+1] = -1
			}
		}
	}
	return pa
}

func (pa TreeAncestor) GetKthAncestor(node, k int) int {
	var m = bits.Len(uint(k))
	for i := 0; i < m; i++ {
		if k>>i&1 > 0 { // k 的二进制从低到高第 i 位是 1
			node = pa[node][i]
			if node < 0 {
				break
			}
		}
	}
	return node
}
