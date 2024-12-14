package main

func continuousSubarrays(nums []int) (ans int64) {
	cnt := make(map[int]int)
	var left int
	for right, num := range nums {
		cnt[num]++
		for {
			mx, mn := num, num
			for k := range cnt {
				mx, mn = max(mx, k), min(mn, k)
			}
			if mx-mn <= 2 {
				break
			}
			x := nums[left]
			if cnt[x]--; cnt[x] == 0 {
				delete(cnt, x)
			}
			left++
		}
		ans += int64(right - left + 1)
	}
	return
}
