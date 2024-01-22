package main

func findErrorNums(nums []int) []int {
	var n, ans = len(nums), make([]int, 2)
	var cnt = make([]bool, n)
	for _, num := range nums {
		if cnt[num-1] {
			ans[0] = num
		} else {
			cnt[num-1] = true
		}
	}
	for i, b := range cnt {
		if !b {
			ans[1] = i + 1
			break
		}
	}
	return ans
}
