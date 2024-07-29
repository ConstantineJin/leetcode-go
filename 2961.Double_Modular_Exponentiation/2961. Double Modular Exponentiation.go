package main

func getGoodIndices(variables [][]int, target int) (ans []int) {
	for i, variable := range variables {
		if pow(pow(variable[0], variable[1], 10), variable[2], variable[3]) == target {
			ans = append(ans, i)
		}
	}
	return
}

func pow(x, n, mod int) int { // 快速幂
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
