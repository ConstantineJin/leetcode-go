package main

func hasTrailingZeros(nums []int) (ans bool) {
	for _, num := range nums {
		if num%2 == 0 {
			if ans {
				return
			} else {
				ans = true
			}
		}
	}
	return false
}
