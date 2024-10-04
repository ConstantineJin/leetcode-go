package main

var diff = []int{1, 2, 2, -1}

func clumsy(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 6
	case 4:
		return 7
	default:
		return n + diff[n%4]
	}
}
