package main

func minimumPossibleSum(n, target int) int {
	var m = min(target/2, n)
	return (m*(m+1) + (target*2+n-m-1)*(n-m)) / 2 % (1e9 + 7)
}
