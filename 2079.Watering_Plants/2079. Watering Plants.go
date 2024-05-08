package main

func wateringPlants(plants []int, capacity int) (ans int) {
	water := capacity
	for i, plant := range plants {
		if water < plant {
			ans += i * 2
			water = capacity
		}
		ans++
		water -= plant
	}
	return
}
