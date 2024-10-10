package main

func numberOfPairs(nums1, nums2 []int, k int) (ans int) {
	mp := make(map[int]int)
	for _, num := range nums1 {
		if num%k > 0 {
			continue
		}
		num /= k
		for d := 1; d*d <= num; d++ {
			if num%d == 0 {
				mp[d]++
				if d*d < num {
					mp[num/d]++
				}
			}
		}
	}
	for _, num := range nums2 {
		ans += mp[num]
	}
	return
}
