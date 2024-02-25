package main

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func resultGrid(image [][]int, threshold int) [][]int {
	var m, n = len(image), len(image[0])
	var res, scale, ok = make([][]int, m), make([][]int, m), make([][][4]bool, m)
	for i, row := range image {
		res[i] = make([]int, n)
		scale[i] = make([]int, n)
		ok[i] = make([][4]bool, n)
		for j, v := range row {
			if i > 0 && abs(v-image[i-1][j]) <= threshold {
				ok[i][j][0] = true // 上
			}
			if j > 0 && abs(v-image[i][j-1]) <= threshold {
				ok[i][j][1] = true // 左
			}
			if i < m-1 && abs(v-image[i+1][j]) <= threshold {
				ok[i][j][2] = true // 下
			}
			if j < n-1 && abs(v-image[i][j+1]) <= threshold {
				ok[i][j][3] = true // 右
			}
		}
	}
	copy(res, image)
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if ok[i][j][0] && ok[i][j][1] && ok[i][j][2] && ok[i][j][3] && ok[i-1][j][1] && ok[i-1][j][3] && ok[i][j-1][0] && ok[i][j-1][2] && ok[i+1][j][1] && ok[i+1][j][3] && ok[i][j+1][0] && ok[i][j+1][2] {
				scale[i][j] = (image[i-1][j-1] + image[i-1][j] + image[i-1][j+1] + image[i][j-1] + image[i][j] + image[i][j+1] + image[i+1][j-1] + image[i+1][j] + image[i+1][j+1]) / 9
			} else {
				scale[i][j] = -1
			}
		}
	}
	for i := range image {
		for j := range image[i] {
			var regCnt int
			var temp int
			if i > 1 && j > 1 && scale[i-1][j-1] >= 0 {
				temp += scale[i-1][j-1]
				regCnt++
			}
			if i > 1 && j > 0 && j < n-1 && scale[i-1][j] >= 0 {
				temp += scale[i-1][j]
				regCnt++
			}
			if i > 1 && j < n-2 && scale[i-1][j+1] >= 0 {
				temp += scale[i-1][j+1]
				regCnt++
			}
			if i > 0 && i < m-1 && j > 1 && scale[i][j-1] >= 0 {
				temp += scale[i][j-1]
				regCnt++
			}
			if i > 0 && j > 0 && i < m-1 && j < n-1 && scale[i][j] >= 0 {
				temp += scale[i][j]
				regCnt++
			}
			if i > 0 && i < m-1 && j < n-2 && scale[i][j+1] >= 0 {
				temp += scale[i][j+1]
				regCnt++
			}
			if i < m-2 && j > 1 && scale[i+1][j-1] >= 0 {
				temp += scale[i+1][j-1]
				regCnt++
			}
			if i < m-2 && j > 0 && j < n-1 && scale[i+1][j] >= 0 {
				temp += scale[i+1][j]
				regCnt++
			}
			if i < m-2 && j < n-2 && scale[i+1][j+1] >= 0 {
				temp += scale[i+1][j+1]
				regCnt++
			}
			if regCnt != 0 {
				res[i][j] = temp / regCnt
			}
		}
	}
	return res
}
