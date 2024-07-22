package main

func losingPlayer(x, y int) string {
	return [2]string{"Bob", "Alice"}[min(x, y/4)%2]
}
