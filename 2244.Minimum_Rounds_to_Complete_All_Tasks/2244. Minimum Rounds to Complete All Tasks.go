package main

func minimumRounds(tasks []int) (ans int) {
	mp := make(map[int]int)
	for _, task := range tasks {
		mp[task]++
	}
	for _, cnt := range mp {
		if cnt < 2 {
			return -1
		}
		ans += (cnt + 2) / 3
	}
	return
}
