package main

func convert(s string, numRows int) string {
	if numRows == 1 || numRows > len(s) {
		return s
	}
	var matrix, c = make([][]byte, numRows), numRows*2 - 2
	for i := 0; i < len(s); i++ {
		var r = i % c
		if r < numRows {
			matrix[r] = append(matrix[r], s[i])
		} else {
			var temp = numRows*2 - r - 2
			matrix[temp] = append(matrix[temp], s[i])
		}
	}
	var ans []byte
	for _, m := range matrix {
		ans = append(ans, m...)
	}
	return string(ans)
}
