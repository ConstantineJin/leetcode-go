package main

func findMaxLength(nums []int) (maxLength int) {
	var mp = map[int]int{0: -1}
	var cnt int
	for i, num := range nums {
		if num == 1 {
			cnt++
		} else {
			cnt--
		}
		if prevIndex, ok := mp[cnt]; ok {
			maxLength = max(maxLength, i-prevIndex)
		} else {
			mp[cnt] = i
		}
	}
	return
}
