package main

import "sort"

func stoneGameVI(aliceValues []int, bobValues []int) int {
	var alice, bob int
	var n = len(aliceValues)
	var values = make([][3]int, n)
	for i := range values {
		values[i] = [3]int{aliceValues[i] + bobValues[i], aliceValues[i], bobValues[i]}
	}
	sort.Slice(values, func(i, j int) bool { return values[i][0] > values[j][0] })
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			alice += values[i][1]
		} else {
			bob += values[i][2]
		}
	}
	if alice > bob {
		return 1
	} else if alice == bob {
		return 0
	} else {
		return -1
	}
}
