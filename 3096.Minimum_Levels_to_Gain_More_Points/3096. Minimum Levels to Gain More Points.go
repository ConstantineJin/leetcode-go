package main

func minimumLevels(possible []int) int {
	n := len(possible)
	prefixSum := make([]int, n+1)
	for i, v := range possible {
		if v == 0 {
			prefixSum[i+1] = prefixSum[i] - 1
		} else {
			prefixSum[i+1] = prefixSum[i] + 1
		}
	}
	for i := 1; i < n; i++ {
		if prefixSum[i]*2 > prefixSum[n] {
			return i
		}
	}
	return -1
}
