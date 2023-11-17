package main

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < (n+1)/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := [4]int{matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i]}
			matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i] = temp[3], temp[0], temp[1], temp[2]
		}
	}
}

func testRotate(matrix [][]int) [][]int {
	rotate(matrix)
	return matrix
}
