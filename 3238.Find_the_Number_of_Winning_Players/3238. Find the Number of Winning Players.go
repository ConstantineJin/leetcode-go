package main

func winningPlayerCount(n int, pick [][]int) (ans int) {
	count := make([][11]int, n)
	for _, p := range pick {
		count[p[0]][p[1]]++
	}
	for i, cnt := range count {
		for _, c := range cnt {
			if c > i {
				ans++
				break
			}
		}
	}
	return
}
