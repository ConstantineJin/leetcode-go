package main

import "math/bits"

func findThePrefixCommonArray(A, B []int) []int {
	ans := make([]int, len(A))
	var p, q uint
	for i := range A {
		p |= 1 << A[i]
		q |= 1 << B[i]
		ans[i] = bits.OnesCount(p & q)
	}
	return ans
}
