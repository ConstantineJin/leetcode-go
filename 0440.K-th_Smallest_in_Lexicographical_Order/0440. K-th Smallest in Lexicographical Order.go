package main

func findKthNumber(n, k int) int {
	getSteps := func(cur int) (res int) {
		first, last := cur, cur
		for first <= n {
			res += min(last, n) - first + 1 // 当前层的节点数
			first *= 10
			last = last*10 + 9
		}
		return
	} // 从当前数字 cur 开始，在以 cur 为前缀的字典树子树下的数字个数
	cur := 1
	k--
	for k > 0 {
		steps := getSteps(cur)
		if steps <= k { // 第 k 小的节点一定不在当前子树中
			k -= steps
			cur++ // 探索下一个相邻的兄弟节点
		} else { // 第 k 小的节点一定在当前子树中
			cur *= 10 // 探索第一个孩子
			k--
		}
	}
	return cur
}
