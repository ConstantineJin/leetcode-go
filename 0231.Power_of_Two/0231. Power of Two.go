package main

func isPowerOfTwo(n int) bool {
	//return bits.OnesCount(uint(n)) == 1
	//return n > 0 && n&(n-1) == 0
	return n > 0 && n&-n == n
}
