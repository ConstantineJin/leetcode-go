package main

//func maxTaxiEarnings(n int, rides [][]int) int64 {
//	sort.Slice(rides, func(i, j int) bool {
//		return rides[i][1] < rides[j][1] // 根据行程终点对rides进行升序排序
//	})
//	var m = len(rides)
//	var dp = make([]int64, m+1) // dp[i+1]表示只接区间[0,i]内的乘客的最大盈利
//	for i := 0; i < m; i++ {
//		var j = sort.Search(i+1, func(k int) bool {
//			return rides[k][1] > rides[i][0]
//		})
//		dp[i+1] = max(dp[i], dp[j]+int64(rides[i][1]-rides[i][0]+rides[i][2])) // 动态规划：对第i位乘客接单：二分查找满足end[j]<=start[i]的最大的j，则dp[i+1]=dp[j]+end[i]-start[i]+tip[i]；对第i位乘客不接单，则dp[i+1]=dp[i]
//	}
//	return dp[m]
//}

func maxTaxiEarnings(n int, rides [][]int) int64 {
	dp := make([]int64, n+1)
	var mp = map[int][][]int{}
	for _, ride := range rides {
		mp[ride[1]] = append(mp[ride[1]], ride)
	}
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1]
		for _, ride := range mp[i] {
			dp[i] = max(dp[i], dp[ride[0]]+int64(ride[1]-ride[0]+ride[2]))
		}
	}
	return dp[n]
}
