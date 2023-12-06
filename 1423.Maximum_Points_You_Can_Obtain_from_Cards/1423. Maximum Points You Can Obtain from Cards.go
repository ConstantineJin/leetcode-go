package main

// 动态规划，超时
//func maxScore(cardPoints []int, k int) int {
//	if k == 0 {
//		return 0
//	}
//	return max(maxScore(cardPoints[:len(cardPoints)-1], k-1)+cardPoints[len(cardPoints)-1], maxScore(cardPoints[1:], k-1)+cardPoints[0])
//}

func maxScore(cardPoints []int, k int) int {
	sum, total, n := 0, 0, len(cardPoints)-k
	for i := 0; i < n; i++ {
		sum += cardPoints[i]
	}
	minSum := sum
	for i := n; i < len(cardPoints); i++ {
		sum += cardPoints[i] - cardPoints[i-n]
		minSum = min(minSum, sum)
	}
	for _, point := range cardPoints {
		total += point
	}
	return total - minSum
}
