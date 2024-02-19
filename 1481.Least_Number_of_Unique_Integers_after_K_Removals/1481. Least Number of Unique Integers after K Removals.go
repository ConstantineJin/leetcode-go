package main

import "sort"

func findLeastNumOfUniqueInts(arr []int, k int) int {
	var mp = make(map[int]int)
	for _, num := range arr {
		mp[num]++
	}
	type pairs struct{ key, value int }
	var s = make([]pairs, 0, len(mp))
	for k, v := range mp {
		s = append(s, pairs{k, v})
	}
	sort.Slice(s, func(i, j int) bool { return s[i].value < s[j].value })
	for len(s) > 0 && k >= s[0].value {
		k -= s[0].value
		s = s[1:]
	}
	return len(s)
}
