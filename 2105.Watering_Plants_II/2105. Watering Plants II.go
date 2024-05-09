package main

func minimumRefill(plants []int, capacityA int, capacityB int) (ans int) {
	i, j := 0, len(plants)-1
	waterA, waterB := capacityA, capacityB
	for i < j {
		if plants[i] > waterA {
			waterA = capacityA
			ans++
		}
		if plants[j] > waterB {
			waterB = capacityB
			ans++
		}
		waterA -= plants[i]
		waterB -= plants[j]
		i++
		j--
	}
	if i == j && max(waterA, waterB) < plants[i] {
		ans++
	}
	return
}
