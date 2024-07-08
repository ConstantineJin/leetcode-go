package main

func maximumPoints(enemyEnergies []int, currentEnergy int) int64 {
	mn := enemyEnergies[0]
	sum := currentEnergy
	for _, energy := range enemyEnergies {
		sum += energy
		mn = min(mn, energy)
	}
	if currentEnergy < mn {
		return 0
	}
	sum -= mn
	return int64(sum / mn)
}
