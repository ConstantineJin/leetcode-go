package main

func countSubarrays(nums []int, minK, maxK int) (ans int64) {
	var minI, maxI, i0 = -1, -1, -1
	for i, x := range nums {
		if x == minK {
			minI = i
		}
		if x == maxK {
			maxI = i
		}
		if x < minK || x > maxK {
			i0 = i // 子数组不能包含 nums[i0]
		}
		ans += int64(max(min(minI, maxI)-i0, 0))
	}
	return
}
