package main

import "sort"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minMovesToSeat(seats []int, students []int) (ans int) {
	sort.Ints(seats)
	sort.Ints(students)
	for i, seat := range seats {
		ans += abs(seat - students[i])
	}
	return
}
