package main

func findKOr(nums []int, k int) (ans int) {
	for i := 0; i < 31; i++ {
		var cnt int
		for _, num := range nums {
			cnt += num >> i & 1
		}
		if cnt >= k {
			ans |= 1 << i
		}
	}
	return
}
