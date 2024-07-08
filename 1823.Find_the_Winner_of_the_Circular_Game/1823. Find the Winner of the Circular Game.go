package main

func findTheWinner(n, k int) int {
	winner := 1
	for i := 2; i <= n; i++ {
		winner = (k+winner-1)%i + 1
	}
	return winner
}
