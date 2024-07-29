package main

func canAliceWin(nums []int) bool {
	var s1, s2 int
	for _, num := range nums {
		if num < 10 {
			s1 += num
		} else {
			s2 += num
		}
	}
	return s1 != s2
}
