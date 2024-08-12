package main

func finalPositionOfSnake(n int, commands []string) (ans int) {
	mp := map[string]int{"UP": -n, "DOWN": n, "LEFT": -1, "RIGHT": 1}
	for _, command := range commands {
		ans += mp[command]
	}
	return
}
