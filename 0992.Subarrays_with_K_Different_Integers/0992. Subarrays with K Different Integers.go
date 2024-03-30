package main

func subarraysWithKDistinct(nums []int, k int) (ans int) {
	var n = len(nums)
	var mp = make(map[int]int)
	for l, r := 0, 0; r < n; r++ {
		mp[nums[r]]++
		if len(mp) >= k {
			ans++
			for ; len(mp) > k; l++ {
				mp[nums[l]]--
				if mp[nums[l]] == 0 {
					delete(mp, nums[l])
				}
			}
			var visit = make(map[int]int)
			for i := l; mp[nums[i]] > visit[nums[i]]+1; i++ {
				visit[nums[i]]++
				ans++
			}
		}
	}
	return
}
