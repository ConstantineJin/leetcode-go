package main

func passThePillow(n int, time int) int {
	time = time%(n*2-2) + 1
	if time <= n {
		return time
	} else {
		return n*2 - time
	}
}
