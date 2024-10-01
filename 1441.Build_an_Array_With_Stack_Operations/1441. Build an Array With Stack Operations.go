package main

func buildArray(target []int, _ int) (ans []string) {
	var prev int
	for _, number := range target {
		for i := 0; i < number-prev-1; i++ {
			ans = append(ans, "Push", "Pop")
		}
		ans = append(ans, "Push")
		prev = number
	}
	return
}
