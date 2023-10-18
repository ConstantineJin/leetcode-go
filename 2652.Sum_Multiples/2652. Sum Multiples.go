package main

func sumOfMultiples(n int) int {
	return ((3+n-n%3)*(n/3) + (5+n-n%5)*(n/5) + (7+n-n%7)*(n/7) - (15+n-n%15)*(n/15) - (21+n-n%21)*(n/21) - (35+n-n%35)*(n/35) + (105+n-n%105)*(n/105)) / 2
}
