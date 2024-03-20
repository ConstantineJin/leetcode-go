package main

const mod int = 1e9 + 7

func minNonZeroProduct(p int) int {
	k := 1<<p - 1
	return k % mod * pow(k-1, p-1) % mod
}

func pow(x, p int) int {
	res := 1
	for x %= mod; p > 0; p-- {
		res = res * x % mod
		x = x * x % mod
	}
	return res
}
