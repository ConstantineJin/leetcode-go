package main

func generate(numRows int) [][]int {
	var ans = make([][]int, numRows)
	ans[0] = []int{1}
	for i := 1; i < numRows; i++ {
		ans[i] = make([]int, i+1)
		ans[i][0], ans[i][i] = 1, 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}
	return ans
}
