package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) (ans []int) {
	mp := make(map[int]int)
	var arr []int
	for _, v := range arr2 {
		mp[v] = 0
	}
	for _, v := range arr1 {
		if _, ok := mp[v]; ok {
			mp[v]++
		} else {
			arr = append(arr, v)
		}
	}
	for _, v := range arr2 {
		for i := 0; i < mp[v]; i++ {
			ans = append(ans, v)
		}
	}
	sort.Ints(arr)
	ans = append(ans, arr...)
	return ans
}
