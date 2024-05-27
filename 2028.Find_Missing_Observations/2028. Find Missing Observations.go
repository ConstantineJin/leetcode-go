package main

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	sum := mean * (m + n)
	ans := make([]int, n)
	var rollsSum int
	for _, roll := range rolls {
		rollsSum += roll
	}
	missingSum := sum - rollsSum
	if missingSum < n || missingSum > n*6 {
		return nil
	}
	average := missingSum / n
	remainder := missingSum - average*n
	for i := range ans {
		ans[i] = average
	}
	for i := 0; i < n && remainder > 0; i++ {
		if remainder > 6-ans[i] {
			remainder -= 6 - ans[i]
			ans[i] = 6
		} else {
			ans[i] += remainder
			remainder = 0
		}
	}
	return ans
}
