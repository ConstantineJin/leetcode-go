package main

func duplicateNumbersXOR(nums []int) (ans int) {
	var mask int
	for _, num := range nums {
		if mask>>num&1 > 0 {
			ans ^= num
		} else {
			mask |= 1 << num
		}
	}
	return
}
