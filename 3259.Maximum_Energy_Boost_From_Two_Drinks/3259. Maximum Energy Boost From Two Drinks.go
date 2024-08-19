package main

func maxEnergyBoost(energyDrinkA, energyDrinkB []int) int64 {
	n := len(energyDrinkA)
	f := make([][2]int, n)
	f[0] = [2]int{energyDrinkA[0], energyDrinkB[0]}
	for i := 1; i < n; i++ {
		f[i] = [2]int{f[i-1][0] + energyDrinkA[i], f[i-1][1] + energyDrinkB[i]}
		if i > 1 {
			f[i][0], f[i][1] = max(f[i][0], f[i-2][1]+energyDrinkA[i]), max(f[i][1], f[i-2][0]+energyDrinkB[i])
		}
	}
	return int64(max(f[n-1][0], f[n-1][1]))
}
