package main

func occurrencesOfElement(nums, queries []int, x int) []int {
	var indices []int
	for i, num := range nums {
		if num == x {
			indices = append(indices, i)
		}
	}
	ans := make([]int, len(queries))
	for i, query := range queries {
		if query <= len(indices) {
			ans[i] = indices[query-1]
		} else {
			ans[i] = -1
		}
	}
	return ans
}
