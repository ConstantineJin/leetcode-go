package main

func distinctDifferenceArray(nums []int) []int {
	var pre, suf, ans = make(map[int]int), make(map[int]int), make([]int, len(nums))
	for _, num := range nums {
		suf[num]++
	}
	for i, num := range nums {
		pre[num]++
		suf[num]--
		if suf[num] == 0 {
			delete(suf, num)
		}
		ans[i] = len(pre) - len(suf)
	}
	return ans
}
