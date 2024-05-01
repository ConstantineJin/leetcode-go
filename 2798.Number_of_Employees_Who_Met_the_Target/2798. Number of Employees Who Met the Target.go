package main

func numberOfEmployeesWhoMetTarget(hours []int, target int) (ans int) {
	for _, hour := range hours {
		if hour >= target {
			ans++
		}
	}
	return
}
