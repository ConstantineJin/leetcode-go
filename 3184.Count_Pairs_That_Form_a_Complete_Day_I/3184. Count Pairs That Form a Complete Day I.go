package main

func countCompleteDayPairs(hours []int) (ans int) {
	n := len(hours)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (hours[i]+hours[j])%24 == 0 {
				ans++
			}
		}
	}
	return
}
