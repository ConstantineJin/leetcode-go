package main

func twoSum(nums []int, target int) []int {
	// map的key为num的值，value为下标
	mp := make(map[int]int)
	for i, num := range nums {
		if _, ok := mp[target-num]; ok {
			return []int{mp[target-num], i}
		} else {
			mp[num] = i
		}
	}
	return make([]int, 0)
}
