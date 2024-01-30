package main

func minimumSeconds(nums []int) int {
	var mp, n = make(map[int][]int), len(nums)
	var res = n
	for i, num := range nums {
		mp[num] = append(mp[num], i)
	}
	for _, pos := range mp {
		var mx = pos[0] + n - pos[len(pos)-1]
		for i := 1; i < len(pos); i++ {
			mx = max(mx, pos[i]-pos[i-1])
		}
		res = min(res, mx/2)
	}
	return res
}
