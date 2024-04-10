package main

func maxSubarrayLength(nums []int, k int) (ans int) {
	var cnt = make(map[int]int)
	var i int
	for j, num := range nums {
		for cnt[num] == k {
			cnt[nums[i]]--
			i++
		}
		cnt[num]++
		ans = max(ans, j-i+1)
	}
	return
}
