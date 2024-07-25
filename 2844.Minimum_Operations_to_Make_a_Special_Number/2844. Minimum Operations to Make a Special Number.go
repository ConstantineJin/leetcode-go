package main

func minimumOperations(num string) (ans int) {
	n := len(num)
	var zeroCnt int
	var meetFive bool
	for i := n - 1; i >= 0; i-- {
		switch num[i] {
		case '0':
			if zeroCnt > 0 { // 后缀可以构成 00
				return n - i - 2
			}
			zeroCnt++
		case '2', '7':
			if meetFive { // 后缀可以构成 25 或 75
				return n - i - 2
			}
		case '5':
			if zeroCnt > 0 { // 后缀可以构成 50
				return n - i - 2
			}
			meetFive = true
		}
	}
	return n - zeroCnt
}
