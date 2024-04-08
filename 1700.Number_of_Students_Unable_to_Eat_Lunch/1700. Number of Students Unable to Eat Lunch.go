package main

func countStudents(students, sandwiches []int) int {
	var one int
	for _, v := range students {
		one += v
	}
	var zero = len(students) - one
	for _, x := range sandwiches {
		if x == 0 && zero > 0 {
			zero--
		} else if x == 1 && one > 0 {
			one--
		} else {
			break
		}
	}
	return zero + one
}
