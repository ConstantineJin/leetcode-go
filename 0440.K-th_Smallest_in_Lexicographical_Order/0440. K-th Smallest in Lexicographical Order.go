package main

func findKthNumber(n, k int) int {
	getSteps := func(cur int) (res int) {
		first, last := cur, cur
		for first <= n {
			res += min(last, n) - first + 1
			first *= 10
			last = last*10 + 9
		}
		return
	}
	cur := 1
	k--
	for k > 0 {
		steps := getSteps(cur)
		if steps <= k { // 第 k 小的节点一定不在当前子树中
			k -= steps
			cur++
		} else { // 第 k 小的节点一定在当前子树中
			cur *= 10
			k--
		}
	}
	return cur
}
