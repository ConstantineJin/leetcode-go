package main

func canBeEqual(target, arr []int) bool {
	mp := make(map[int]int)
	for _, v := range target {
		mp[v]++
	}
	for _, v := range arr {
		if mp[v] == 0 {
			return false
		}
		mp[v]--
	}
	return true
}
