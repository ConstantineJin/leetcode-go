package main

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	if tomatoSlices%2 == 1 || tomatoSlices/2-cheeseSlices < 0 || cheeseSlices*2-tomatoSlices/2 < 0 {
		return make([]int, 0)
	}
	return []int{tomatoSlices/2 - cheeseSlices, cheeseSlices*2 - tomatoSlices/2}
}
