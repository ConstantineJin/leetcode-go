package main

const mx = 1e5 + 1

var isPrime [mx]bool

func init() { // 线性筛求质数
	var primes []int
	for i := 2; i < mx; i++ {
		isPrime[i] = true
	}
	for i := 2; i < mx; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
		for _, prime := range primes {
			if i*prime >= mx {
				break
			}
			isPrime[i*prime] = false
			if i%prime == 0 {
				break
			}
		}
	}
}

func countPaths(n int, edges [][]int) int64 {
	var g = make([][]int, n+1)
	for _, edge := range edges {
		var i, j = edge[0], edge[1]
		g[i] = append(g[i], j)
		g[j] = append(g[j], i)
	}

	var dfs func(int, int)
	var seen []int
	dfs = func(i, pre int) {
		seen = append(seen, i)
		for _, j := range g[i] {
			if j != pre && !isPrime[j] {
				dfs(j, i)
			}
		}
	}
	var ans int
	var count = make([]int, n+1)
	for i := 1; i <= n; i++ {
		if !isPrime[i] { // 从质数节点开始DFS
			continue
		}
		var cur int              // 记录目前为止所经过的所有以i为根的子树的大小之和
		for _, j := range g[i] { // 遍历i的邻接节点
			if isPrime[j] { // 遇到下一个质数节点结束
				continue
			}
			if count[j] == 0 {
				seen = []int{} // 每一次dfs所经过的节点集合
				dfs(j, 0)
				var cnt = len(seen) // 以当前质数为根，合法路径构成的子树大小
				for _, k := range seen {
					count[k] = cnt
				}
			}
			ans += count[j] * cur
			cur += count[j]
		}
		ans += cur
	}
	return int64(ans)
}
