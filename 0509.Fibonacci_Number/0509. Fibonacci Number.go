package main

import "math"

// 递归
//func fib(n int) int {
//	switch n {
//	case 0, 1:
//		return n
//	default:
//		return fib(n-2) + fib(n-1)
//	}
//}

// 递推
//func fib(n int) int {
//	if n < 2 {
//		return n
//	}
//	p, q, r := 0, 0, 1
//	for i := 2; i <= n; i++ {
//		p, q = q, r
//		r = p + q
//	}
//	return r
//}

// 通项公式
func fib(n int) int {
	sqrt5 := math.Sqrt(5)
	return int(math.Round((math.Pow((1+sqrt5)/2, float64(n)) - math.Pow((1-sqrt5)/2, float64(n))) / sqrt5))
}
