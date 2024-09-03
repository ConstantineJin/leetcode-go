package main

func countOfPairs(n, x, y int) []int64 {
	x, y = min(x, y), max(x, y) // 始终保证 x 不在 y 的右侧
	ans := make([]int64, n)
	if x+1 >= y { // 额外边是自环或者连接相邻两个房屋，不会产生任何效用
		for i := 1; i < n; i++ {
			ans[i-1] = int64(n-i) * 2
		}
		return ans
	}
	diff := make([]int, n+1) // 差分数组
	add := func(l, r int) { diff[l]++; diff[r+1]-- }
	for i := 1; i < n; i++ {
		if i <= x { // 如果 i 不在 x 的右侧
			k := (x + y + 1) / 2  // 分界点
			add(1, k-i)           // 对于 [i+1,k] 内的房子，都可以直接到达
			add(x-i+2, x-i+y-k)   // 对于 [k+1,y-1] 内的房子，可以用额外边到达
			add(x-i+1, x-i+1+n-y) // 对于 [y,n] 内的房子，可以用额外边到达
		} else if i < (x+y)/2 { // 如果 i 在 x 到 x 和 y 的中点之间
			k := i + (y-x+1)/2    // 分界点
			add(1, k-i)           // 对于 [i+1,k] 内的房子，都可以直接到达
			add(i-x+2, i-x+y-k)   // 对于 [k+1,y-1] 内的房子，可以用额外边到达
			add(i-x+1, i-x+1+n-y) // 对于 [y,n] 内的房子，可以用额外边到达
		} else { // 对于更大的 i，无法利用额外边缩短它到右侧房屋的距离，把区间 [1,n-i] 都加上 1
			add(1, n-i)
		}
	}
	var sumD int64
	for i, d := range diff[1:] {
		sumD += int64(d)
		ans[i] = sumD * 2 // 根据题意 (i,j) 和 (j,i) 算两对
	}
	return ans
}
