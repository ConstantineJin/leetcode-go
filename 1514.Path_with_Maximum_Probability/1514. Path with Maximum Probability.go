package main

import "container/heap"

type pair struct {
	to          int
	probability float64 // 从起点到当前节点的累积概率
}

func maxProbability(n int, edges [][]int, successProbability []float64, startNode, endNode int) float64 {
	g := make([][]pair, n)
	for i, edge := range edges {
		u, v, prob := edge[0], edge[1], successProbability[i]
		g[u] = append(g[u], pair{v, prob})
		g[v] = append(g[v], pair{u, prob})
	}
	probabilities := make([]float64, n)
	q := &hp{}
	heap.Init(q)
	heap.Push(q, pair{to: startNode, probability: 1})
	probabilities[startNode] = 1
	for q.Len() > 0 { // Dijkstra 算法 + 优先队列
		cur := heap.Pop(q).(pair)
		if cur.to == endNode { // 由于使用大顶堆实现，因此每个节点第一次被访问到时就是当前可以到达该节点的最大概率
			return cur.probability
		}
		if cur.probability < probabilities[cur.to] {
			continue
		}
		for _, p := range g[cur.to] {
			to, prob := p.to, p.probability
			newProb := prob * cur.probability
			if newProb > probabilities[to] { // 松弛操作
				probabilities[to] = newProb
				heap.Push(q, pair{to: to, probability: newProb})
			}
		} // 完成循环后 cur.to 也被加入了 S 集合
	}
	return 0 // startNode 与 endNode 不联通
}

type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].probability > h[j].probability } // 大顶堆
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
