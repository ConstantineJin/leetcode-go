package main

func busyStudent(startTime, endTime []int, queryTime int) (ans int) {
	for i := range startTime {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			ans++
		}
	}
	return
}
