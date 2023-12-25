package main

func maximumSumOfHeights(maxHeights []int) int64 {
	var ans, n = 0, len(maxHeights)
	var suf, st, sum = make([]int, n+1), []int{n}, 0
	for i := n - 1; i >= 0; i-- {
		for len(st) > 1 && maxHeights[i] <= maxHeights[st[len(st)-1]] {
			var j = st[len(st)-1]
			st = st[:len(st)-1]
			sum -= maxHeights[j] * (st[len(st)-1] - j)
		}
		sum += maxHeights[i] * (st[len(st)-1] - i)
		suf[i] = sum
		st = append(st, i)
	}
	ans = sum
	st = []int{-1}
	var pre int
	for i, height := range maxHeights {
		for len(st) > 1 && height <= maxHeights[st[len(st)-1]] {
			var j = st[len(st)-1]
			st = st[:len(st)-1]
			pre -= maxHeights[j] * (j - st[len(st)-1])
		}
		pre += height * (i - st[len(st)-1])
		ans = max(ans, pre+suf[i+1])
		st = append(st, i)
	}
	return int64(ans)
}
