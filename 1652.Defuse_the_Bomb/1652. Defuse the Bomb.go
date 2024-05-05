package main

func decrypt(code []int, k int) []int {
	n := len(code)
	ans := make([]int, n)
	if k > 0 {
		for i := range code {
			for j := 1; j <= k; j++ {
				ans[i] += code[(i+j)%n]
			}
		}
	} else if k < 0 {
		for i := range code {
			for j := -1; j >= k; j-- {
				ans[i] += code[(i+j+n)%n]
			}
		}
	}
	return ans
}
