package main

func xorQueries(arr []int, queries [][]int) []int {
	n := len(arr)
	prefixXor := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixXor[i+1] = prefixXor[i] ^ arr[i]
	}
	ans := make([]int, 0, len(queries))
	for _, query := range queries {
		ans = append(ans, prefixXor[query[0]]^prefixXor[query[1]+1])
	}
	return ans
}
