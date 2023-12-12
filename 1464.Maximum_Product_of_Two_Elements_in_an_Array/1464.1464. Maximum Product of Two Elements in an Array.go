package main

func maxProduct(nums []int) int {
	var a, b int
	for _, num := range nums {
		a, b = max(a, num), max(b, min(a, num))
	}
	return (a - 1) * (b - 1)
}
