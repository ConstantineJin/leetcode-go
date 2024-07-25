package main

func minimumOperations(num string) (ans int) {
	n := len(num)
	var meetZero, meetFive bool
	for i := n - 1; i >= 0; i-- {
		switch num[i] {
		case '0':
			if meetZero { // 后缀可以构成 00
				return n - i - 2
			}
			meetZero = true
		case '2', '7':
			if meetFive { // 后缀可以构成 25 或 75
				return n - i - 2
			}
		case '5':
			if meetZero { // 后缀可以构成 50
				return n - i - 2
			}
			meetFive = true
		}
	}
	if meetZero {
		return n - 1
	}
	return n
}
