package main

func xorN(n int) int {
	switch n % 4 {
	case 0:
		return n
	case 1:
		return 1
	case 2:
		return n + 1
	default:
		return 0
	}
}

func xorOperation(n, start int) int {
	a := start / 2
	b := n & start & 1 // 都为奇数才是 1
	return (xorN(a+n-1)^xorN(a-1))*2 + b
}
