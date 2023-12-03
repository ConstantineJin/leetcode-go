package main

func tribonacci(n int) int {
	switch n {
	case 0, 1:
		return n

	case 2:
		return 1
	default:
		p, q, r, s := 0, 0, 1, 1
		for i := 3; i <= n; i++ {
			p, q, r = q, r, s
			s = p + q + r
		}
		return s
	}
}
