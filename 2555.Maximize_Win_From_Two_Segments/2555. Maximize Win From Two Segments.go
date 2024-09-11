package main

func maximizeWin(prizePositions []int, k int) (ans int) {
	n := len(prizePositions)
	if prizePositions[n-1]-prizePositions[0] <= k*2+1 { // 如果两段线段可以完全覆盖所有奖品
		return n
	}
	var mx, left, right int                     // mx 表示第一条线段所能覆盖的最多奖品个数
	for mid, position := range prizePositions { // 枚举 mid
		for right < n && prizePositions[right]-position <= k { // 先将 mid 作为第二条线段的左端点，计算其能覆盖到的最右端下标
			right++
		}
		ans = max(ans, mx+right-mid)
		for position-prizePositions[left] > k { // 再将 mid 作为第一条线段的右端点，计算其能覆盖到的最左端下标
			left++
		}
		mx = max(mx, mid-left+1)
	}
	return
}
