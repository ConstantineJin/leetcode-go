package main

func containsNearbyDuplicate(nums []int, k int) bool {
	var mp = make(map[int][]int)
	for i, num := range nums {
		mp[num] = append(mp[num], i)
	}
	for _, arr := range mp {
		for i := 1; i < len(arr); i++ {
			if arr[i]-arr[i-1] <= k {
				return true
			}
		}
	}
	return false
}
