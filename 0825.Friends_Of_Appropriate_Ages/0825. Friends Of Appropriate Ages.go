package main

func numFriendRequests(ages []int) (ans int) {
	var cnt [121]int
	for _, age := range ages {
		cnt[age]++
	}
	var window, ageY int
	for ageX, c := range cnt {
		window += c
		if ageY*2 <= ageX+14 {
			window -= cnt[ageY]
			ageY++
		}
		if window > 0 {
			ans += c*window - c
		}
	}
	return
}
