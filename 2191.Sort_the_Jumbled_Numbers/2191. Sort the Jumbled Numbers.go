package main

import "sort"

type pair struct{ original, new int }

func sortJumbled(mapping []int, nums []int) []int {
	pairs := make([]pair, len(nums))
	for i, num := range nums {
		var arr []int
		if num == 0 {
			arr = []int{mapping[0]}
		}
		for num > 0 {
			arr = append(arr, mapping[num%10])
			num /= 10
		}
		for j := len(arr) - 1; j >= 0; j-- {
			num = num*10 + arr[j]
		}
		pairs[i] = pair{original: nums[i], new: num}
	}
	sort.SliceStable(pairs, func(i, j int) bool { return pairs[i].new < pairs[j].new })
	for i, p := range pairs {
		nums[i] = p.original
	}
	return nums
}
