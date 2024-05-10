package main

func countTestedDevices(batteryPercentages []int) (ans int) {
	for _, percentage := range batteryPercentages {
		if percentage > ans {
			ans++
		}
	}
	return
}
