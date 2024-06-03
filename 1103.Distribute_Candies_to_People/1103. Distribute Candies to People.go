package main

import "math"

// 方法1:模拟
//func distributeCandies(candies, numPeople int) []int {
//	ans := make([]int, numPeople)
//	for i := 1; candies > 0; i++ {
//		ans[(i-1)%numPeople] += min(candies, i)
//		candies -= i
//	}
//	return ans
//}

// 方法2:数学
func distributeCandies(candies, n int) []int {
	m := int((math.Sqrt(float64(8*candies+1)) - 1) / 2)
	k, extra := m/n, m%n
	ans := make([]int, n)
	for i := range ans {
		j := k
		if i < extra {
			j++
		}
		ans[i] = j*(j-1)/2*n + j*(i+1)
	}
	ans[extra] += candies - m*(m+1)/2
	return ans
}
