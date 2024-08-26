package main

func maxOperations(nums []int, k int) (ans int) {
	cnt := make(map[int]int)
	for _, num := range nums {
		if cnt[k-num] > 0 {
			cnt[k-num]--
			ans++
		} else {
			cnt[num]++
		}
	}
	return
}
