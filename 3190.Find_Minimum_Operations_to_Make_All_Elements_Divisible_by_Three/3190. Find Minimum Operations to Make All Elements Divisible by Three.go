package main

func minimumOperations(nums []int) (ans int) {
	for _, num := range nums {
		if num%3 != 0 {
			ans++
		}
	}
	return
}
