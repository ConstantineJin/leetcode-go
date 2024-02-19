package main

func rearrangeArray(nums []int) []int {
	var n = len(nums) / 2
	var pos, neg, ans = make([]int, n), make([]int, n), make([]int, n*2)
	var i, j int
	for _, num := range nums {
		if num > 0 {
			pos[i] = num
			i++
		} else {
			neg[j] = num
			j++
		}
	}
	for i = 0; i < n; i++ {
		ans[i*2] = pos[i]
		ans[i*2+1] = neg[i]
	}
	return ans
}
