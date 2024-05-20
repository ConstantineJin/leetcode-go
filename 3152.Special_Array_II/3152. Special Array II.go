package main

func isArraySpecial(nums []int, queries [][]int) []bool {
	m, n := len(queries), len(nums)
	pos := make([]int, n)
	for i := 1; i < len(nums); i++ {
		if nums[i-1]%2 == nums[i]%2 {
			pos[i] = 1
		}
	}
	sum := make([]int, n+1)
	for i, po := range pos {
		sum[i+1] = sum[i] + po
	}
	ans := make([]bool, m)
	for i, query := range queries {
		ans[i] = sum[query[0]+1] == sum[query[1]+1]
	}
	return ans
}
