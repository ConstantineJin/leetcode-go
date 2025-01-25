package main

func minimumMoney(transactions [][]int) int64 {
	var totalLose, mx int
	for _, transaction := range transactions {
		cost, cashBack := transaction[0], transaction[1]
		totalLose += max(cost-cashBack, 0)
		mx = max(mx, min(cost, cashBack))
	}
	return int64(totalLose + mx)
}
