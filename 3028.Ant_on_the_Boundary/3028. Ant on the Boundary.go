package main

func returnToBoundaryCount(nums []int) (ans int) {
	var pos int
	for _, num := range nums {
		pos += num
		if pos == 0 {
			ans++
		}
	}
	return
}
